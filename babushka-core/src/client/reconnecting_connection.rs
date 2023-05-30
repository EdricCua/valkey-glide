use crate::connection_request::{AddressInfo, TlsMode};
use crate::retry_strategies::RetryStrategy;
use futures_intrusive::sync::ManualResetEvent;
use logger_core::{log_debug, log_warn};
use redis::aio::{ConnectionLike, MultiplexedConnection};
use redis::{RedisConnectionInfo, RedisError, RedisResult};
use std::sync::Arc;
use std::time::Duration;
use tokio::sync::Mutex;
use tokio::task;
use tokio_retry::Retry;

use super::get_connection_info;

/// The object that is used in order to recreate a connection after a disconnect.
struct ConnectionBackend {
    /// This signal is reset when a connection disconnects, and set when a new `ConnectionState` has been set with either a `Connected` or a `Disconnected` state.
    /// Clone of the connection who experience the disconnect can wait on the signal in order to be notified when the new connection state is established.
    connection_available_signal: ManualResetEvent,
    /// Information needed in order to create a new connection.
    connection_info: redis::Client,
}

/// State of the current connection. Allows the user to use a connection only when a reconnect isn't in progress or has failed.
enum ConnectionState {
    /// A connection has been made, and hasn't disconnected yet.
    Connected(MultiplexedConnection, Arc<ConnectionBackend>),
    /// There's a reconnection effort on the way, no need to try reconnecting again.
    Reconnecting(Arc<ConnectionBackend>),
}

/// This allows us to safely share and replace the connection state between clones of the client.
type StateWrapper = Arc<Mutex<ConnectionState>>;

#[derive(Clone)]
pub(super) struct ReconnectingConnection {
    state: StateWrapper,
}

async fn try_create_connection(
    connection_backend: Arc<ConnectionBackend>,
    retry_strategy: RetryStrategy,
) -> RedisResult<StateWrapper> {
    let client = &connection_backend.connection_info;
    let action = || {
        log_warn(
            // TODO -log_debug
            "connection creation",
            format!("Creating multiplexed connection"),
        );
        client.get_multiplexed_async_connection()
    };

    let connection = Retry::spawn(retry_strategy.get_iterator(), action).await?;
    Ok(Arc::new(Mutex::new(ConnectionState::Connected(
        connection,
        connection_backend,
    ))))
}

fn get_client(
    address: &AddressInfo,
    tls_mode: TlsMode,
    redis_connection_info: redis::RedisConnectionInfo,
) -> RedisResult<redis::Client> {
    redis::Client::open(get_connection_info(
        address,
        tls_mode,
        redis_connection_info,
    ))
}

/// This iterator isn't exposed to users, and can't be configured.
fn internal_retry_iterator() -> impl Iterator<Item = Duration> {
    const MAX_DURATION: Duration = Duration::from_secs(5);
    crate::retry_strategies::get_exponential_backoff(
        crate::retry_strategies::EXPONENT_BASE,
        crate::retry_strategies::FACTOR,
        crate::retry_strategies::NUMBER_OF_RETRIES,
    )
    .get_iterator()
    .chain(std::iter::repeat(MAX_DURATION))
}

impl ReconnectingConnection {
    pub(super) async fn new(
        address: &AddressInfo,
        connection_retry_strategy: RetryStrategy,
        redis_connection_info: RedisConnectionInfo,
        tls_mode: TlsMode,
    ) -> RedisResult<Self> {
        log_debug(
            "connection creation",
            format!("Attempting connection to {address}"),
        );

        let client = Arc::new(ConnectionBackend {
            connection_info: get_client(address, tls_mode, redis_connection_info)?,
            connection_available_signal: ManualResetEvent::new(true),
        });
        let state = try_create_connection(client, connection_retry_strategy).await?;
        log_debug(
            "connection creation",
            format!("Connection to {address} created"),
        );
        Ok(Self { state })
    }

    async fn get_connection(&self) -> Result<MultiplexedConnection, RedisError> {
        loop {
            // Using a limited scope in order to release the mutex lock before waiting for notifications.
            let backend = {
                let mut guard = self.state.lock().await;
                match &mut *guard {
                    ConnectionState::Reconnecting(backend) => backend.clone(),
                    ConnectionState::Connected(connection, _) => {
                        return Ok(connection.clone());
                    }
                }
            };
            backend.connection_available_signal.wait().await;
        }
    }

    async fn reconnect(&self) {
        let backend = {
            let mut guard = self.state.lock().await;
            let backend = match &*guard {
                ConnectionState::Connected(_, backend) => {
                    backend.connection_available_signal.reset();
                    backend.clone()
                }
                _ => {
                    log_warn(
                        // TODO -log_trace
                        "reconnect",
                        format!("already started"),
                    );
                    // exit early - if reconnection already started or failed, there's nothing else to do.
                    return;
                }
            };
            *guard = ConnectionState::Reconnecting(backend.clone());
            backend
        };
        log_warn(
            // TODO -log_debug
            "reconnect",
            format!("starting"),
        );
        let clone = self.clone();
        // The reconnect task is spawned instead of awaited here, so that if this task will be dropped for some reason, the reconnection attempt will continue.
        task::spawn(async move {
            let client = &backend.connection_info;
            for sleep_duration in internal_retry_iterator() {
                log_warn(
                    // TODO -log_debug
                    "connection creation",
                    format!("Creating multiplexed connection"),
                );
                match client.get_multiplexed_async_connection().await {
                    Ok(connection) => {
                        let mut guard = clone.state.lock().await;
                        log_warn(
                            // TODO -log_debug
                            "reconnect",
                            format!("completed succesfully"),
                        );
                        backend.connection_available_signal.set();
                        *guard = ConnectionState::Connected(connection, backend);
                        break;
                    }
                    Err(_) => tokio::time::sleep(sleep_duration).await,
                }
            }
        });
    }

    pub(super) async fn send_packed_command(
        &mut self,
        cmd: &redis::Cmd,
    ) -> redis::RedisResult<redis::Value> {
        log_warn(
            // TODO -log_trace
            "ReconnectingConnection",
            format!("sending command"),
        );
        let mut connection = self.get_connection().await?;
        let result = connection.send_packed_command(cmd).await;
        match result {
            Err(err) if err.is_connection_dropped() => {
                self.reconnect().await;
                Err(err)
            }
            _ => result,
        }
    }

    pub(super) async fn send_packed_commands(
        &mut self,
        cmd: &redis::Pipeline,
        offset: usize,
        count: usize,
    ) -> redis::RedisResult<Vec<redis::Value>> {
        let mut connection = self.get_connection().await?;
        let result = connection.send_packed_commands(cmd, offset, count).await;
        match result {
            Err(err) if err.is_connection_dropped() => {
                self.reconnect().await;
                Err(err)
            }
            _ => result,
        }
    }

    pub(super) fn get_db(&self) -> i64 {
        let guard = self.state.blocking_lock();
        match &*guard {
            ConnectionState::Connected(connection, _) => connection.get_db(),
            _ => -1,
        }
    }
}
