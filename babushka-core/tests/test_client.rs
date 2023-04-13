mod utilities;

#[cfg(test)]
mod client {
    use super::*;
    use babushka::client::ClientCMD;
    use babushka::connection_request::{ConnectionRequest, ConnectionRetryStrategy};
    use protobuf::MessageField;
    use redis::aio::ConnectionLike;
    use redis::{RedisResult, Value};
    use rstest::rstest;

    use std::time::Duration;

    use utilities::*;

    struct TestBasics {
        server: RedisServer,
        client: ClientCMD,
    }

    async fn setup_test_basics_and_connection_retry_strategy(
        use_tls: bool,
        connection_retry_strategy: Option<ConnectionRetryStrategy>,
    ) -> TestBasics {
        let server = TestContext::new(ServerType::Tcp { tls: use_tls }).server;
        let address_info = get_address_info(server.get_client_addr());
        let mut connection_request = ConnectionRequest::new();
        connection_request.addresses = vec![address_info];
        connection_request.use_tls = use_tls;
        connection_request.connection_retry_strategy =
            MessageField::from_option(connection_retry_strategy);
        let client = ClientCMD::create_client(connection_request).await.unwrap();
        TestBasics { server, client }
    }

    async fn setup_test_basics(use_tls: bool) -> TestBasics {
        setup_test_basics_and_connection_retry_strategy(use_tls, None).await
    }

    async fn send_get(client: &mut ClientCMD, key: &str) -> RedisResult<Value> {
        let mut get_command = redis::Cmd::new();
        get_command.arg("GET").arg(key);
        client.req_packed_command(&get_command).await
    }

    async fn send_set_and_get(mut client: ClientCMD, key: String) {
        const VALUE_LENGTH: usize = 10;
        let value = generate_random_string(VALUE_LENGTH);

        let mut set_command = redis::Cmd::new();
        set_command.arg("SET").arg(key.as_str()).arg(value.clone());
        let set_result = client.req_packed_command(&set_command).await.unwrap();
        let get_result = send_get(&mut client, key.as_str()).await.unwrap();

        assert_eq!(set_result, Value::Okay);
        assert_eq!(get_result, Value::Data(value.into_bytes()));
    }

    #[rstest]
    #[timeout(Duration::from_millis(10000))]
    fn test_send_set_and_get(#[values(false, true)] use_tls: bool) {
        block_on_all(async {
            let test_basics = setup_test_basics(use_tls).await;
            let key = "hello";
            send_set_and_get(test_basics.client.clone(), key.to_string()).await;
        });
    }

    #[rstest]
    #[timeout(Duration::from_millis(15000))]
    fn test_client_handle_concurrent_workload(#[values(false, true)] use_tls: bool) {
        block_on_all(async {
            let test_basics = setup_test_basics(use_tls).await;
            const VALUE_LENGTH: usize = 1000000;
            const NUMBER_OF_CONCURRENT_OPERATIONS: usize = 1000;

            let mut actions = Vec::with_capacity(NUMBER_OF_CONCURRENT_OPERATIONS);
            for index in 0..NUMBER_OF_CONCURRENT_OPERATIONS {
                actions.push(send_set_and_get(
                    test_basics.client.clone(),
                    format!("key{index}"),
                ));
            }
            futures::future::join_all(actions).await;
        });
    }

    #[rstest]
    #[timeout(Duration::from_millis(10000))]
    fn test_report_closing_when_server_closes() {
        block_on_all(async {
            let mut test_basics = setup_test_basics(false).await;
            let server = test_basics.server;
            drop(server);

            let get_result = send_get(&mut test_basics.client, "foobar")
                .await
                .unwrap_err();
            assert!(get_result.is_connection_dropped());
        });
    }

    #[rstest]
    #[timeout(Duration::from_millis(10000))]
    fn test_reconnect_and_complete_request_after_temporary_disconnect(
        #[values(false, true)] use_tls: bool,
    ) {
        block_on_all(async move {
            // We want a retry strategy that retries often, so that the test won't hang an wait too long for a retry.
            let mut retry_strategy = ConnectionRetryStrategy::new();
            retry_strategy.number_of_retries = 100;
            retry_strategy.factor = 1;
            retry_strategy.exponent_base = 2;
            let mut test_basics =
                setup_test_basics_and_connection_retry_strategy(use_tls, Some(retry_strategy))
                    .await;
            let server = test_basics.server;
            let address = server.get_client_addr().clone();
            drop(server);

            let thread = std::thread::spawn(move || {
                block_on_all(async move {
                    let get_result = send_get(
                        &mut test_basics.client,
                        "test_reconnect_and_complete_request_after_temporary_disconnect",
                    )
                    .await
                    .unwrap();
                    assert_eq!(get_result, Value::Nil);
                });
            });
            let _new_server = RedisServer::new_with_addr_and_modules(address.clone(), &[]);
            wait_for_server_to_become_ready(&address).await;

            thread.join().unwrap();
        });
    }
}
