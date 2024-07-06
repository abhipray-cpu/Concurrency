package main

import (
	"log"
	"net"
	"net/rpc"
)

func StartServer() {
	arithmeticService := new(ArithMeticService)
	rpc.Register(arithmeticService)

	rpcService := new(RPCService)
	rpc.Register(rpcService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	defer listener.Close()
	log.Println("Server listening on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
