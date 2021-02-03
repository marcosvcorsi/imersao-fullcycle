package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:#{port}")

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start grpc server", err)
	}
}
