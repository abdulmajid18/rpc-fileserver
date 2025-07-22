package demo

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func RpcClient() {
	var serverAddress = "localhost"
	conn, err := net.Dial("tcp", serverAddress+":1234")

	if err != nil {
		log.Fatal("Dailing RPC server failed", err)
	}

	client := rpc.NewClient(conn)
	args := &Args{7, 8}

	var reply int
	err = client.Call("Arith.Multiple", args, &reply)

	if err != nil {
		log.Fatal("Client couldn't process response", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
