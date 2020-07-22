package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/leopedroso45/codeedu/framework/pb"

	"github.com/leopedroso45/codeedu/application/repositories"
	"github.com/leopedroso45/codeedu/application/usecases"

	"github.com/leopedroso45/codeedu/framework/servers"

	"github.com/jinzhu/gorm"
	"github.com/leopedroso45/codeedu/framework/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var db *gorm.DB

func main() {

	db = utils.ConnectDB()
	db.LogMode(true)

	port := flag.Int("port", 0, "Choose the server port:")
	flag.Parse()
	log.Printf("Starting server at port: %d", *port)

	userServer := setUpUserServer()
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot start a server: ", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start a server: ", err)
	}

}

func setUpUserServer() *servers.UserServer {
	userRepository := repositories.UserRepositoryDb{Db: db}
	userServer := servers.NewUserServer()
	userServer.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}
	return userServer
}
