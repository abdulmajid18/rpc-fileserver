package rpcclient

import (
	"fmt"
	"net"
	"net/rpc"
)

func Call(method string, address string, args any, reply any) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to dial RPC server: %w", err)
	}
	defer conn.Close()
	client := rpc.NewClient(conn)
	newMethod := fmt.Sprintf("%s%s", "FileOperations.", method)
	err = client.Call(newMethod, args, reply)
	if err != nil {
		return fmt.Errorf("RPC call failed: %w", err)
	}
	return nil
}
