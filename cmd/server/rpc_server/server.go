package rpcserver

import (
	"context"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type RPCServer struct {
	listener net.Listener
	wg       sync.WaitGroup
	shutdown chan struct{}
	rpc      *rpc.Server
}

func NewRPCServer(addr string) (*RPCServer, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	server := &RPCServer{
		listener: ln,
		shutdown: make(chan struct{}),
		rpc:      rpc.NewServer(),
	}
	return server, nil
}

func (s *RPCServer) RegisterService(service any) *RPCServer {
	if err := rpc.Register(service); err != nil {
		log.Fatalf("failed to register service: %v", err)
	}
	return s
}

func (s *RPCServer) Start() {
	log.Printf("RPC server starting on %s", s.listener.Addr().String())
	go s.handleShutdown()
	for {
		select {
		case <-s.shutdown:
			log.Println("Server is shutting down")
			return
		default:
			if tcpListener, ok := s.listener.(*net.TCPListener); ok {
				tcpListener.SetDeadline(time.Now().Add(1 * time.Second))
			}

			conn, err := s.listener.Accept()
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}

			}
			if conn == nil {
				log.Println("Received nil connection, continuing...")
				continue
			}
			log.Printf("New client connected from %s", conn.RemoteAddr().String())
			s.wg.Add(1)
			go s.handleConnection(conn)

		}
	}
}

func (s *RPCServer) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer func() {
		conn.Close()
		log.Printf("Client %s disconnected", conn.RemoteAddr().String())
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-s.shutdown:
			conn.Close()
		case <-ctx.Done():
		}
	}()

	rpc.ServeConn(conn)

}

func (s *RPCServer) handleShutdown() {
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	<-quitCh
	log.Println("Shutdown signal received")

	close(s.shutdown)

	if err := s.listener.Close(); err != nil {
		log.Printf("Error closing listener: %v", err)
	}

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("All connections closed gracefully")
	case <-time.After(10 * time.Second):
		log.Println("Timeout waiting for connections to close")
	}

	log.Println("RPC server shut down gracefully")
}
