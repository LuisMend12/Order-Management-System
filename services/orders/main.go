package main

import (
	"log"
	"os"
)

func main() {
	addr := os.Getenv("ORDERS_GRPC_ADDR")
	if addr == "" {
		addr = ":50051"
	}

	grpcServer := NewGRPCServer(addr)
	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}

}
