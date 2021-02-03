package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/mvcorsi/imersao-fullcycle/codepix-go/application/grpc/pb"
	"github.com/mvcorsi/imersao-fullcycle/codepix-go/application/usecases"
	"github.com/mvcorsi/imersao-fullcycle/codepix-go/infra/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repositories.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecases.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)

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
