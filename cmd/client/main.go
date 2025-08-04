package main

import (
	rpcclient "abdulmajid/fileserver/cmd/client/rpc_client"
	"fmt"
	"time"
)

func main() {
	// Wait a bit for server to start if running locally
	time.Sleep(1 * time.Second)
	fmt.Println("Starting RPC client...")
	rpcclient.RpcClient()
}
