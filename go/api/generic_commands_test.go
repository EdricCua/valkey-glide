// Copyright Valkey GLIDE Project Contributors - SPDX Identifier: Apache-2.0

package api

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/valkey-io/valkey-glide/go/api/options"
)

func ExampleGlideClient_CustomCommand() {
	var client *GlideClient = getExampleGlideClient() // example helper function
	result, err := client.CustomCommand(context.Background(), []string{"ping"})
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}
	fmt.Println(result)

	// Output: PONG
}

func ExampleGlideClient_Move() {
	var client *GlideClient = getExampleGlideClient() // example helper function
	key := uuid.New().String()
	_, err := client.Set(context.Background(), key, "hello")
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}
	result, err := client.Move(context.Background(), key, 2)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}
	fmt.Println(result)

	// Output: true
}

func ExampleGlideClient_Scan() {
	var client *GlideClient = getExampleGlideClient() // example helper function
	client.CustomCommand(context.Background(), []string{"FLUSHALL"})
	client.Set(context.Background(), "key1", "hello")
	resCursor, resCollection, err := client.Scan(context.Background(), 0)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}
	fmt.Println("Cursor:", resCursor)
	fmt.Println("Collection:", resCollection)

	// Output:
	// Cursor: 0
	// Collection: [key1]
}

func ExampleGlideClient_ScanWithOptions() {
	var client *GlideClient = getExampleGlideClient() // example helper function
	opts := options.NewScanOptions().SetCount(10).SetType(options.ObjectTypeList)
	client.CustomCommand(context.Background(), []string{"FLUSHALL"})
	client.LPush(context.Background(), "key1", []string{"1", "3", "2", "4"})
	resCursor, resCollection, err := client.ScanWithOptions(context.Background(), 0, *opts)
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}
	fmt.Println("Cursor:", resCursor)
	fmt.Println("Collection:", resCollection)

	// Output:
	// Cursor: 0
	// Collection: [key1]
}

func ExampleGlideClient_RandomKey() {
	var client *GlideClient = getExampleGlideClient() // example helper function
	key := uuid.New().String()
	client.Set(context.Background(), key, "Hello")
	result, err := client.RandomKey(context.Background())
	if err != nil {
		fmt.Println("Glide example failed with an error: ", err)
	}
	fmt.Println(len(result.Value()) > 0)

	// Output: true
}
