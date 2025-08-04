package rpcclient

import (
	"abdulmajid/fileserver/cmd/client/request"
	"abdulmajid/fileserver/cmd/client/response"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func RpcClient() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing RPC server failed:", err)
	}
	defer conn.Close()

	client := rpc.NewClient(conn)

	args := &request.DirRequest{Name: "new_dirr"}
	var reply response.GenericResponse

	err = client.Call("FileOperations.CreateDir", args, &reply)
	if err != nil {
		log.Fatal("RPC call failed:", err)
	}

	if reply.Success {
		fmt.Println("✅ Success:", reply.Message)
	} else {
		fmt.Println("❌ Failed:", reply.Message)
	}
}
