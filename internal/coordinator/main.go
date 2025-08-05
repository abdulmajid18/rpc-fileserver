package main

import (
	"abdulmajid/fileserver/internal/coordinator/etcd"
	"log"
)

func main() {
	client := etcd.EtcdClient{}
	err := client.InitClient()
	if err != nil {
		log.Fatal("Failed to initialize etcd client:", err)
	}
	defer client.Close()

	instanceID := etcd.GetEnv("INSTANCE_ID", "rpc1")
	err = client.RegisterService(instanceID, "0.0.0.0:5001")
	if err != nil {
		log.Fatal("failed to register service: ", err)
	}
}
