package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/mvcorsi/imersao-fullcycle/codepix-go/application/grpc"
	"github.com/mvcorsi/imersao-fullcycle/codepix-go/infra/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
}
