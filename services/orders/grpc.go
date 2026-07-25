package main

import (
	"fmt"
	"log"
	"net"

	handler "github.com/ArdiDev1/Order-Management-System/services/orders/handler/orders"
	"github.com/ArdiDev1/Order-Management-System/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.addr, err)
	}

	grpcServer := grpc.NewServer()

	// register grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Printf("Starting gRPC server on port :50051")

	return grpcServer.Serve(lis)
}
