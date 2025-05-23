// Copyright Valkey GLIDE Project Contributors - SPDX Identifier: Apache-2.0

package api

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/valkey-io/valkey-glide/go/api/config"
	"github.com/valkey-io/valkey-glide/go/api/options"
)

var (
	libraryCode = `#!lua name=mylib
redis.register_function{ function_name = 'myfunc', callback = function(keys, args) return 42 end, flags = { 'no-writes' } }`
	libraryCodeWithArgs = `#!lua name=mylib
redis.register_function{ function_name = 'myfunc', callback = function(keys, args) return args[1] end, flags = { 'no-writes' } }`
)

// FunctionLoad Examples
func ExampleGlideClient_FunctionLoad() {
	client := getExampleGlideClient()

	result, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// mylib
}

func ExampleGlideClusterClient_FunctionLoad() {
	client := getExampleGlideClusterClient()

	result, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// mylib
}

func ExampleGlideClusterClient_FunctionLoadWithRoute() {
	client := getExampleGlideClusterClient()

	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	result, err := client.FunctionLoadWithRoute(context.Background(), libraryCode, true, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// mylib
}

// FunctionFlush Examples
func ExampleGlideClient_FunctionFlush() {
	client := getExampleGlideClient()

	result, err := client.FunctionFlush(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClusterClient_FunctionFlush() {
	client := getExampleGlideClusterClient()

	result, err := client.FunctionFlush(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClusterClient_FunctionFlushWithRoute() {
	client := getExampleGlideClusterClient()

	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	result, err := client.FunctionFlushWithRoute(context.Background(), opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClient_FunctionFlushSync() {
	client := getExampleGlideClient()

	result, err := client.FunctionFlushSync(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClusterClient_FunctionFlushSync() {
	client := getExampleGlideClient()

	result, err := client.FunctionFlushSync(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClusterClient_FunctionFlushSyncWithRoute() {
	client := getExampleGlideClusterClient()

	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	result, err := client.FunctionFlushSyncWithRoute(context.Background(), opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClient_FunctionFlushAsync() {
	client := getExampleGlideClient()

	result, err := client.FunctionFlushAsync(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClusterClient_FunctionFlushAsync() {
	client := getExampleGlideClusterClient()

	result, err := client.FunctionFlushAsync(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)
}

func ExampleGlideClusterClient_FunctionFlushAsyncWithRoute() {
	client := getExampleGlideClusterClient()

	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	result, err := client.FunctionFlushAsyncWithRoute(context.Background(), opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

// FCall Examples
func ExampleGlideClient_FCall() {
	client := getExampleGlideClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	fcallResult, err := client.FCall(context.Background(), "myfunc")
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(fcallResult)

	// Output:
	// 42
}

func ExampleGlideClusterClient_FCall() {
	client := getExampleGlideClusterClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	fcallResult, err := client.FCall(context.Background(), "myfunc")
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(fcallResult)

	// Output:
	// 42
}

func ExampleGlideClusterClient_FCallWithRoute() {
	client := getExampleGlideClusterClient()

	// Load function
	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	_, err := client.FunctionLoadWithRoute(context.Background(), libraryCode, true, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	result, err := client.FCallWithRoute(context.Background(), "myfunc", opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	for _, value := range result.MultiValue() {
		fmt.Println(value)
		break
	}

	// Output:
	// 42
}

func ExampleGlideClient_FCallWithKeysAndArgs() {
	client := getExampleGlideClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCodeWithArgs, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	key1 := "{testKey}-" + uuid.New().String()
	key2 := "{testKey}-" + uuid.New().String()
	result, err := client.FCallWithKeysAndArgs(context.Background(), "myfunc", []string{key1, key2}, []string{"3", "4"})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// 3
}

func ExampleGlideClusterClient_FCallWithKeysAndArgs() {
	client := getExampleGlideClusterClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCodeWithArgs, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	key1 := "{testKey}-" + uuid.New().String()
	key2 := "{testKey}-" + uuid.New().String()
	result, err := client.FCallWithKeysAndArgs(context.Background(), "myfunc", []string{key1, key2}, []string{"3", "4"})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// 3
}

func ExampleGlideClusterClient_FCallWithArgs() {
	client := getExampleGlideClusterClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCodeWithArgs, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	result, err := client.FCallWithArgs(context.Background(), "myfunc", []string{"1", "2"})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result.SingleValue())

	// Output:
	// 1
}

func ExampleGlideClusterClient_FCallWithArgsWithRoute() {
	client := getExampleGlideClusterClient()

	// Load function
	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	_, err := client.FunctionLoadWithRoute(context.Background(), libraryCodeWithArgs, true, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	result, err := client.FCallWithArgsWithRoute(context.Background(), "myfunc", []string{"1", "2"}, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	for _, value := range result.MultiValue() {
		fmt.Println(value)
		break
	}

	// Output:
	// 1
}

func ExampleGlideClient_FCallReadOnly() {
	client := getExampleGlideClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	fcallResult, err := client.FCallReadOnly(context.Background(), "myfunc")
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(fcallResult)

	// Output:
	// 42
}

func ExampleGlideClusterClient_FCallReadOnly() {
	client := getExampleGlideClusterClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	fcallResult, err := client.FCallReadOnly(context.Background(), "myfunc")
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(fcallResult)

	// Output:
	// 42
}

func ExampleGlideClusterClient_FCallReadOnlyWithRoute() {
	client := getExampleGlideClusterClient()

	// Load function
	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	_, err := client.FunctionLoadWithRoute(context.Background(), libraryCode, true, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	result, err := client.FCallReadOnlyWithRoute(context.Background(), "myfunc", opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	for _, value := range result.MultiValue() {
		fmt.Println(value)
		break
	}

	// Output:
	// 42
}

func ExampleGlideClient_FCallReadOnlyWithKeysAndArgs() {
	client := getExampleGlideClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCodeWithArgs, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	key1 := "{testKey}-" + uuid.New().String()
	key2 := "{testKey}-" + uuid.New().String()
	result, err := client.FCallReadOnlyWithKeysAndArgs(
		context.Background(),
		"myfunc",
		[]string{key1, key2},
		[]string{"3", "4"},
	)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// 3
}

func ExampleGlideClusterClient_FCallReadOnlyWithKeysAndArgs() {
	client := getExampleGlideClusterClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCodeWithArgs, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	key1 := "{testKey}-" + uuid.New().String()
	key2 := "{testKey}-" + uuid.New().String()
	result, err := client.FCallReadOnlyWithKeysAndArgs(
		context.Background(),
		"myfunc",
		[]string{key1, key2},
		[]string{"3", "4"},
	)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// 3
}

func ExampleGlideClusterClient_FCallReadOnlyWithArgs() {
	client := getExampleGlideClusterClient()

	// Load function
	_, err := client.FunctionLoad(context.Background(), libraryCodeWithArgs, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	result, err := client.FCallReadOnlyWithArgs(context.Background(), "myfunc", []string{"1", "2"})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result.SingleValue())

	// Output:
	// 1
}

func ExampleGlideClusterClient_FCallReadOnlyWithArgsWithRoute() {
	client := getExampleGlideClusterClient()

	// Load function
	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	_, err := client.FunctionLoadWithRoute(context.Background(), libraryCodeWithArgs, true, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	// Call function
	result, err := client.FCallReadOnlyWithArgsWithRoute(context.Background(), "myfunc", []string{"1", "2"}, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	for _, value := range result.MultiValue() {
		fmt.Println(value)
		break
	}

	// Output:
	// 1
}

func ExampleGlideClient_FunctionStats() {
	client := getExampleGlideClient()

	// Load a function first
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Get function statistics
	stats, err := client.FunctionStats(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Print statistics for each node
	for _, nodeStats := range stats {
		fmt.Println("Example stats:")
		for engineName, engine := range nodeStats.Engines {
			fmt.Printf("  Engine %s: %d functions, %d libraries\n",
				engineName, engine.FunctionCount, engine.LibraryCount)
		}
	}

	// Output:
	// Example stats:
	//   Engine LUA: 1 functions, 1 libraries
}

func ExampleGlideClusterClient_FunctionStats() {
	client := getExampleGlideClusterClient()

	// Load a function first
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Get function statistics
	stats, err := client.FunctionStats(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Print statistics
	fmt.Printf("Nodes reached: %d\n", len(stats))
	for _, nodeStats := range stats {
		fmt.Println("Example stats:")
		for engineName, engine := range nodeStats.Engines {
			fmt.Printf("  Engine %s: %d functions, %d libraries\n",
				engineName, engine.FunctionCount, engine.LibraryCount)
		}
		break
	}

	// Output:
	// Nodes reached: 6
	// Example stats:
	//   Engine LUA: 1 functions, 1 libraries
}

func ExampleGlideClusterClient_FunctionStatsWithRoute() {
	client := getExampleGlideClusterClient()

	// Load a function first
	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	_, err := client.FunctionLoadWithRoute(context.Background(), libraryCode, true, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Get function statistics with route
	stats, err := client.FunctionStatsWithRoute(context.Background(), opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Print statistics
	fmt.Printf("Nodes reached: %d\n", len(stats.MultiValue()))
	for _, nodeStats := range stats.MultiValue() {
		fmt.Println("Example stats:")
		for engineName, engine := range nodeStats.Engines {
			fmt.Printf("  Engine %s: %d functions, %d libraries\n",
				engineName, engine.FunctionCount, engine.LibraryCount)
		}
		break
	}

	// Output:
	// Nodes reached: 3
	// Example stats:
	//   Engine LUA: 1 functions, 1 libraries
}

func ExampleGlideClient_FunctionDelete() {
	client := getExampleGlideClient()

	// Load a function first
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Delete function
	result, err := client.FunctionDelete(context.Background(), "mylib")
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClusterClient_FunctionDelete() {
	client := getExampleGlideClusterClient()

	// Load a function first
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Delete function
	result, err := client.FunctionDelete(context.Background(), "mylib")
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClusterClient_FunctionDeleteWithRoute() {
	client := getExampleGlideClusterClient()

	// Load a function first
	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	_, err := client.FunctionLoadWithRoute(context.Background(), libraryCode, true, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Delete function with route
	result, err := client.FunctionDeleteWithRoute(context.Background(), "mylib", opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Println(result)

	// Output:
	// OK
}

func ExampleGlideClient_FunctionKill() {
	client := getExampleGlideClient()

	// Try to kill when no function is running
	_, err := client.FunctionKill(context.Background())
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// Output:
	// Expected error: An error was signalled by the server: - NotBusy: No scripts in execution right now.
}

func ExampleGlideClusterClient_FunctionKill() {
	client := getExampleGlideClusterClient()

	// Try to kill when no function is running
	_, err := client.FunctionKill(context.Background())
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// Output:
	// Expected error: An error was signalled by the server: - NotBusy: No scripts in execution right now.
}

func ExampleGlideClusterClient_FunctionKillWithRoute() {
	client := getExampleGlideClusterClient()

	// Try to kill with route when no function is running
	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	_, err := client.FunctionKillWithRoute(context.Background(), opts)
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// Output:
	// Expected error: An error was signalled by the server: - NotBusy: No scripts in execution right now.
}

func ExampleGlideClient_FunctionList() {
	client := getExampleGlideClient()

	// Load a function first
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	query := FunctionListQuery{
		LibraryName: "mylib",
		WithCode:    true,
	}

	libs, err := client.FunctionList(context.Background(), query)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Printf("There are %d libraries loaded.\n", len(libs))
	for i, lib := range libs {
		fmt.Printf("%d) Library name '%s', on engine %s, with %d functions\n", i+1, lib.Name, lib.Engine, len(lib.Functions))
		for j, fn := range lib.Functions {
			fmt.Printf("   %d) function '%s'\n", j+1, fn.Name)
		}
	}
	// Output:
	// There are 1 libraries loaded.
	// 1) Library name 'mylib', on engine LUA, with 1 functions
	//    1) function 'myfunc'
}

func ExampleGlideClusterClient_FunctionList() {
	client := getExampleGlideClusterClient()

	// Load a function first
	_, err := client.FunctionLoad(context.Background(), libraryCode, true)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	query := FunctionListQuery{
		LibraryName: "mylib",
		WithCode:    true,
	}

	libs, err := client.FunctionList(context.Background(), query)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}

	fmt.Printf("There are %d libraries loaded.\n", len(libs))
	for i, lib := range libs {
		fmt.Printf("%d) Library name '%s', on engine %s, with %d functions\n", i+1, lib.Name, lib.Engine, len(lib.Functions))
		for j, fn := range lib.Functions {
			fmt.Printf("   %d) function '%s'\n", j+1, fn.Name)
		}
	}
	// Output:
	// There are 1 libraries loaded.
	// 1) Library name 'mylib', on engine LUA, with 1 functions
	//    1) function 'myfunc'
}

func ExampleGlideClusterClient_FunctionListWithRoute() {
	client := getExampleGlideClusterClient()

	// Load a function first
	route := config.Route(config.AllPrimaries)
	opts := options.RouteOption{
		Route: route,
	}
	_, err := client.FunctionLoadWithRoute(context.Background(), libraryCode, true, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// List functions with route
	query := FunctionListQuery{
		WithCode: true,
	}
	result, err := client.FunctionListWithRoute(context.Background(), query, opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Print results for each node
	for _, libs := range result.MultiValue() {
		fmt.Println("Example Node:")
		for _, lib := range libs {
			fmt.Printf("  Library: %s\n", lib.Name)
			fmt.Printf("  Engine: %s\n", lib.Engine)
			fmt.Printf("  Functions: %d\n", len(lib.Functions))
		}
		break
	}

	// Output:
	// Example Node:
	//   Library: mylib
	//   Engine: LUA
	//   Functions: 1
}

func ExampleGlideClient_FunctionDump() {
	client := getExampleGlideClient()

	// Call FunctionDump to get the serialized payload of all loaded libraries
	dump, _ := client.FunctionDump(context.Background())
	if len(dump) > 0 {
		fmt.Println("Function dump got a payload")
	}

	// Output:
	// Function dump got a payload
}

func ExampleGlideClusterClient_FunctionDump() {
	client := getExampleGlideClusterClient()

	// Call FunctionDump to get the serialized payload of all loaded libraries
	dump, _ := client.FunctionDump(context.Background())
	if len(dump) > 0 {
		fmt.Println("Function dump got a payload")
	}

	// Output:
	// Function dump got a payload
}

func ExampleGlideClusterClient_FunctionDumpWithRoute() {
	client := getExampleGlideClusterClient()

	// Call FunctionDumpWithRoute to get the serialized payload of all loaded libraries with a route
	dump, _ := client.FunctionDumpWithRoute(context.Background(), config.RandomRoute)
	if len(dump.SingleValue()) > 0 {
		fmt.Println("Function dump got a payload")
	}

	// Output:
	// Function dump got a payload
}

func ExampleGlideClient_FunctionRestore() {
	client := getExampleGlideClient()

	// Attempt to restore with invalid dump data
	invalidDump := "invalid_dump_data"
	_, err := client.FunctionRestore(context.Background(), invalidDump)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Output:
	// Error: An error was signalled by the server: - ResponseError: DUMP payload version or checksum are wrong
}

func ExampleGlideClusterClient_FunctionRestore() {
	client := getExampleGlideClusterClient()

	// Attempt to restore with invalid dump data
	invalidDump := "invalid_dump_data"
	_, err := client.FunctionRestore(context.Background(), invalidDump)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Output:
	// Error: An error was signalled by the server: - ResponseError: DUMP payload version or checksum are wrong
}

func ExampleGlideClusterClient_FunctionRestoreWithRoute() {
	client := getExampleGlideClusterClient()

	// Attempt to restore with invalid dump data and route
	invalidDump := "invalid_dump_data"
	route := config.RandomRoute
	_, err := client.FunctionRestoreWithRoute(context.Background(), invalidDump, route)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Output:
	// Error: An error was signalled by the server: - ResponseError: DUMP payload version or checksum are wrong
}

func ExampleGlideClient_FunctionRestoreWithPolicy() {
	client := getExampleGlideClient()

	// Attempt to restore with invalid dump data and policy
	invalidDump := "invalid_dump_data"
	_, err := client.FunctionRestoreWithPolicy(context.Background(), invalidDump, options.FlushPolicy)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Output:
	// Error: An error was signalled by the server: - ResponseError: DUMP payload version or checksum are wrong
}

func ExampleGlideClusterClient_FunctionRestoreWithPolicy() {
	client := getExampleGlideClusterClient()

	// Attempt to restore with invalid dump data and policy
	invalidDump := "invalid_dump_data"
	_, err := client.FunctionRestoreWithPolicy(context.Background(), invalidDump, options.FlushPolicy)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Output:
	// Error: An error was signalled by the server: - ResponseError: DUMP payload version or checksum are wrong
}

func ExampleGlideClusterClient_FunctionRestoreWithPolicyWithRoute() {
	client := getExampleGlideClusterClient()

	// Attempt to restore with invalid dump data, policy and route
	invalidDump := "invalid_dump_data"
	route := config.RandomRoute
	_, err := client.FunctionRestoreWithPolicyWithRoute(context.Background(), invalidDump, options.FlushPolicy, route)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	// Output:
	// Error: An error was signalled by the server: - ResponseError: DUMP payload version or checksum are wrong
}

func ExampleGlideClient_InvokeScript() {
	client := getExampleGlideClient()

	// Create a simple Lua script that returns a string
	script := options.NewScript("return 'Hello from Lua'")

	// Execute the script
	result, err := client.InvokeScript(context.Background(), *script)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	fmt.Println(result)

	// Output:
	// Hello from Lua
}

func ExampleGlideClusterClient_InvokeScript() {
	client := getExampleGlideClusterClient()

	// Create a simple Lua script that returns a number
	script := options.NewScript("return 123")

	// Execute the script
	result, err := client.InvokeScript(context.Background(), *script)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	fmt.Println(result)

	// Output:
	// 123
}

func ExampleGlideClient_InvokeScriptWithOptions() {
	client := getExampleGlideClient()

	// Create a Lua script that uses keys and arguments
	scriptText := `
		local key = KEYS[1]
		local value = ARGV[1]
		redis.call('SET', key, value)
		return redis.call('GET', key)
	`
	script := options.NewScript(scriptText)

	// Create a unique key for testing
	testKey := "test-key-" + uuid.New().String()

	// Set up script options with keys and arguments
	scriptOptions := options.NewScriptOptions().
		WithKeys([]string{testKey}).
		WithArgs([]string{"Hello World"})

	// Execute the script with options
	result, err := client.InvokeScriptWithOptions(context.Background(), *script, *scriptOptions)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	fmt.Println(result)

	// Output:
	// Hello World
}

func ExampleGlideClusterClient_InvokeScriptWithOptions() {
	client := getExampleGlideClusterClient()

	// Create a Lua script that performs calculations with arguments
	scriptText := `
		local a = tonumber(ARGV[1])
		local b = tonumber(ARGV[2])
		return a + b
	`
	script := options.NewScript(scriptText)

	// Set up script options with arguments
	scriptOptions := options.NewScriptOptions().
		WithArgs([]string{"10", "20"})

	// Execute the script with options
	result, err := client.InvokeScriptWithOptions(context.Background(), *script, *scriptOptions)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	fmt.Println(result)

	// Output:
	// 30
}

func ExampleGlideClusterClient_InvokeScriptWithClusterOptions() {
	client := getExampleGlideClusterClient()

	// Create a Lua script.
	scriptText := "return 'Hello'"

	script := options.NewScript(scriptText)

	// Set up cluster script options
	clusterScriptOptions := options.NewClusterScriptOptions()

	// Set the route
	clusterScriptOptions.Route = config.AllPrimaries

	// Execute the script with cluster options
	result, err := client.InvokeScriptWithClusterOptions(context.Background(), *script, *clusterScriptOptions)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Print the result. The result contains response from multiple nodes.
	// We are checking and printing the response from only one node below.
	for _, value := range result.MultiValue() {
		if value != nil && value.(string) == "Hello" {
			fmt.Println(value)
			break
		}
	}

	// Output:
	// Hello
}

func ExampleGlideClient_ScriptExists() {
	client := getExampleGlideClient()

	// Invoke a script
	script := options.NewScript("return 'Hello World!'")
	client.InvokeScript(context.Background(), *script)

	response, err := client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println(response)

	// Cleanup
	script.Close()

	// Output: [true]
}

func ExampleGlideClusterClient_ScriptExists() {
	client := getExampleGlideClusterClient()

	// Invoke a script
	script := options.NewScript("return 'Hello World!'")
	client.InvokeScript(context.Background(), *script)

	response, err := client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println(response)

	// Cleanup
	script.Close()

	// Output: [true]
}

func ExampleGlideClusterClient_ScriptExistsWithRoute() {
	client := getExampleGlideClusterClient()
	route := options.RouteOption{Route: config.NewSlotKeyRoute(config.SlotTypePrimary, "1")}

	// Invoke a script
	script := options.NewScript("return 'Hello World!'")
	client.InvokeScriptWithRoute(context.Background(), *script, route)

	response, err := client.ScriptExistsWithRoute(context.Background(), []string{script.GetHash()}, route)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println(response)

	// Cleanup
	script.Close()

	// Output: [true]
}

func ExampleGlideClient_ScriptFlush() {
	client := getExampleGlideClient()

	// First, load a script
	script := options.NewScript("return 'Hello World!'")
	_, err := client.InvokeScript(context.Background(), *script)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Verify script exists
	exists, err := client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists before flush:", exists[0])

	// Flush all scripts
	result, err := client.ScriptFlush(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Flush result:", result)

	// Verify script no longer exists
	exists, err = client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists after flush:", exists[0])

	// Cleanup
	script.Close()

	// Output:
	// Script exists before flush: true
	// Flush result: OK
	// Script exists after flush: false
}

func ExampleGlideClient_ScriptFlushWithMode() {
	client := getExampleGlideClient()

	// First, load a script
	script := options.NewScript("return 'Hello World!'")
	_, err := client.InvokeScript(context.Background(), *script)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Verify script exists
	exists, err := client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists before flush:", exists[0])

	// Flush all scripts with ASYNC mode
	result, err := client.ScriptFlushWithMode(context.Background(), options.ASYNC)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Flush result:", result)

	// Verify script no longer exists
	exists, err = client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists after flush:", exists[0])

	// Cleanup
	script.Close()

	// Output:
	// Script exists before flush: true
	// Flush result: OK
	// Script exists after flush: false
}

func ExampleGlideClusterClient_ScriptFlush() {
	client := getExampleGlideClusterClient()

	// First, load a script
	script := options.NewScript("return 'Hello World!'")
	_, err := client.InvokeScript(context.Background(), *script)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Verify script exists
	exists, err := client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists before flush:", exists[0])

	// Flush all scripts
	result, err := client.ScriptFlush(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Flush result:", result)

	// Verify script no longer exists
	exists, err = client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists after flush:", exists[0])

	// Cleanup
	script.Close()

	// Output:
	// Script exists before flush: true
	// Flush result: OK
	// Script exists after flush: false
}

func ExampleGlideClusterClient_ScriptFlushWithMode() {
	client := getExampleGlideClusterClient()

	// First, load a script
	script := options.NewScript("return 'Hello World!'")
	_, err := client.InvokeScript(context.Background(), *script)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Verify script exists
	exists, err := client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists before flush:", exists[0])

	// Flush all scripts with ASYNC mode
	result, err := client.ScriptFlushWithMode(context.Background(), options.ASYNC)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Flush result:", result)

	// Verify script no longer exists
	exists, err = client.ScriptExists(context.Background(), []string{script.GetHash()})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists after flush:", exists[0])

	// Cleanup
	script.Close()

	// Output:
	// Script exists before flush: true
	// Flush result: OK
	// Script exists after flush: false
}

func ExampleGlideClusterClient_ScriptFlushWithOptions() {
	client := getExampleGlideClusterClient()
	route := options.RouteOption{Route: config.AllPrimaries}

	// First, load a script on all primaries
	script := options.NewScript("return 'Hello World!'")
	_, err := client.InvokeScriptWithRoute(context.Background(), *script, route)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Verify script exists
	exists, err := client.ScriptExistsWithRoute(context.Background(), []string{script.GetHash()}, route)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists before flush:", exists[0])

	// Flush all scripts on all primaries with ASYNC mode
	scriptFlushOptions := options.NewScriptFlushOptions().WithMode(options.ASYNC).WithRoute(&route)
	result, err := client.ScriptFlushWithOptions(context.Background(), *scriptFlushOptions)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Flush result:", result)

	// Verify script no longer exists
	exists, err = client.ScriptExistsWithRoute(context.Background(), []string{script.GetHash()}, route)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}
	fmt.Println("Script exists after flush:", exists[0])

	// Cleanup
	script.Close()

	// Output:
	// Script exists before flush: true
	// Flush result: OK
	// Script exists after flush: false
}

func ExampleGlideClient_ScriptKill() {
	client := getExampleGlideClient()

	// Try to kill scripts when no scripts are running
	_, err := client.ScriptKill(context.Background())
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// Output:
	// Expected error: An error was signalled by the server: - NotBusy: No scripts in execution right now.
}

func ExampleGlideClusterClient_ScriptKill_withoutRoute() {
	client := getExampleGlideClusterClient()

	// Try to kill scripts when no scripts are running
	_, err := client.ScriptKill(context.Background())
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// Output:
	// Expected error: An error was signalled by the server: - NotBusy: No scripts in execution right now.
}

func ExampleGlideClusterClient_ScriptKill_withRoute() {
	key := "{randomkey}1"
	client := getExampleGlideClusterClient()

	// Create a route with our specified key
	route := options.RouteOption{
		Route: config.NewSlotKeyRoute(config.SlotTypePrimary, key),
	}

	// Try to kill scripts when no scripts are running
	_, err := client.ScriptKillWithRoute(context.Background(), route)
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// Output:
	// Expected error: An error was signalled by the server: - NotBusy: No scripts in execution right now.
}

// ScriptShow Examples
func ExampleGlideClient_ScriptShow() {
	client := getExampleGlideClient()

	// First, create and invoke a script
	scriptText := "return 'Hello, World!'"
	script := options.NewScript(scriptText)
	_, err := client.InvokeScript(context.Background(), *script)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Now show the script source using ScriptShow
	scriptSource, err := client.ScriptShow(context.Background(), script.GetHash())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	fmt.Println(scriptSource)

	// Output:
	// return 'Hello, World!'
}

func ExampleGlideClusterClient_ScriptShow() {
	client := getExampleGlideClusterClient()

	// First, create and invoke a script
	scriptText := "return 'Hello World'"
	script := options.NewScript(scriptText)
	_, err := client.InvokeScript(context.Background(), *script)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	// Now show the script source using ScriptShow
	scriptSource, err := client.ScriptShow(context.Background(), script.GetHash())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
		return
	}

	fmt.Println(scriptSource)

	// Output:
	// return 'Hello World'
}
