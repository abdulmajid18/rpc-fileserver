package main

import (
	rpcserver "abdulmajid/fileserver/cmd/server/rpc_server"
	"abdulmajid/fileserver/internal/fileservice"
	"log"
)

func main() {
	// go demo.RpcServer()
	var addr string = "localhost:1234"
	server, err := rpcserver.NewRPCServer(addr)

	if err != nil {
		log.Fatalf("Failed to create RPC server: %v", err)
	}

	server.RegisterService(new(fileservice.FileOperations))

	server.Start()
}
