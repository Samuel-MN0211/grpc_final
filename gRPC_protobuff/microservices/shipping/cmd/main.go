package main

import (
	"log"
	"net"

	"github.com/ruandg/microservices-proto/golang/shipping/microservices-proto/shipping"
	grpc_adapter "github.com/ruandg/microservices/shipping/internal/adapters/grpc"
	api_pkg "github.com/ruandg/microservices/shipping/internal/application/core/api"
	"google.golang.org/grpc"
)

func main() {
	api := api_pkg.NewAPI()
	server := grpc_adapter.NewServer(api)
	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, server)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Shipping microservice running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
