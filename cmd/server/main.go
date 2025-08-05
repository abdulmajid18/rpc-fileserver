package main

import (
	rpcserver "abdulmajid/fileserver/cmd/server/rpc_server"
	"abdulmajid/fileserver/internal/fileservice"
	"log"
	"os"
)

func main() {
	// go demo.RpcServer()
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}
	addr := ":" + port
	server, err := rpcserver.NewRPCServer(addr)

	if err != nil {
		log.Fatalf("Failed to create RPC server: %v", err)
	}

	server.RegisterService(new(fileservice.FileOperations))

	server.Start()
}
