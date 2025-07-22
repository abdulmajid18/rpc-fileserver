package main

import (
	"abdulmajid/fileserver/cmd/server/demo"
	rpcserver "abdulmajid/fileserver/cmd/server/rpc_server"
	"log"
)

func main() {
	// go demo.RpcServer()
	var addr string = "localhost:1234"
	server, err := rpcserver.NewRPCServer(addr)

	if err != nil {
		log.Fatalf("Failed to create RPC server: %v", err)
	}

	server.RegisterService(new(demo.Arith))

	server.Start()
}
