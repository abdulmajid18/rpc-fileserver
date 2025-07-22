package demo

import (
	"log"
	"net"
	"net/rpc"
)

func RpcServer() {
	arith := new(Arith)
	rpc.Register(arith)
	ln, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("Server cannot listen to connection ....", err)
	}

	log.Println("RPC server waiting for request ...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Listener cannot accept incomming condition")
		}
		go rpc.ServeConn(conn)

	}
}
