package main

import (
	"abdulmajid/fileserver/cmd/server/demo"
	"fmt"
	"time"
)

func main() {
	// Wait a bit for server to start if running locally
	time.Sleep(1 * time.Second)
	fmt.Println("Starting RPC client...")
	demo.RpcClient()
}
