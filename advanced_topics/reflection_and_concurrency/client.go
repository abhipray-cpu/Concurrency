package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func RunClient() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	// Call ArithmeticService.Add
	args := AddArgs{A: 10, B: 20}
	var reply int
	err = client.Call("ArithMeticService.Add", args, &reply)
	if err != nil {
		log.Fatal("ArithMeticService error:", err)
	}
	fmt.Printf("ArithMeticService.Add: %d + %d = %d\n", args.A, args.B, reply)

	// Call RPCService.Concat
	concatArgs := ConcatArgs{Str1: "Hello, ", Str2: "World!"}
	var concatReply string
	err = client.Call("RPCService.Concat", concatArgs, &concatReply)
	if err != nil {
		log.Fatal("RPCService error:", err)
	}
	fmt.Printf("RPCService.Concat: %s\n", concatReply)
}
