package main

import (
	"abdulmajid/fileserver/cmd/server/demo"
	"time"
)

func main() {
	go demo.RpcServer()

	time.Sleep(5 * time.Second)

	demo.RpcClient()
}
