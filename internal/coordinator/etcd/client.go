package etcd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdClient struct {
	Client *clientv3.Client
}

func (etcdClient *EtcdClient) InitClient() error {
	ectcdEndpoint := os.Getenv("ETCD_ENDPOINTS")
	endpoints := []string{"http://etcd1:2379", "http://etcd2:2379", "http://localhost:2379", "http://0.0.0.0:2379", ectcdEndpoint}

	cfg := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to etcd: %w", err)
	}

	etcdClient.Client = client
	return nil
}

func (etcdClient EtcdClient) RegisterService(instanceID, addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lease, err := etcdClient.Client.Grant(ctx, 10)
	log.Printf("Lease created for: %s \n", instanceID)
	if err != nil {
		return fmt.Errorf("failed to create lease: %w", err)
	}

	// Register with a key like "/services/rpc_servers/{instanceID}"
	key := "/services/rpc_servers/" + instanceID
	value := addr // e.g., "10.0.0.1:50051"

	_, err = etcdClient.Client.Put(ctx, key, value, clientv3.WithLease(lease.ID))
	log.Printf("Etcd client created with key and value : %s \n", key)
	if err != nil {
		return fmt.Errorf("failed to register service: %w", err)
	}

	log.Printf("Service registered: %s -> %s", key, addr)

	// Keep lease alive in background
	go func() {
		keepAlive, err := etcdClient.Client.KeepAlive(ctx, lease.ID)
		if err != nil {
			log.Printf("Failed to start keepalive: %v", err)
			return
		}

		for {
			select {
			case ka, ok := <-keepAlive:
				if !ok {
					log.Println("Keepalive channel closed")
					return
				}
				log.Printf("Lease renewed: ID=%d TTL=%d", ka.ID, ka.TTL)
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (etcdClient *EtcdClient) Close() {
	if etcdClient.Client != nil {
		etcdClient.Client.Close()
	}
}

func (etcdClient EtcdClient) ListServers() (map[string]string, error) {
	resp, err := etcdClient.Client.Get(context.Background(), "/services/rpc_servers/", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	servers := make(map[string]string)
	for _, kv := range resp.Kvs {
		instanceID := string(kv.Key[len("/services/rpc_servers/"):])
		servers[instanceID] = string(kv.Value) // Map instanceID → address
	}
	return servers, nil
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
