package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/johnsmera/chall/application/repositories"
	"github.com/johnsmera/chall/application/repositories/usecases"
	"github.com/johnsmera/chall/infra/db"
	"github.com/johnsmera/chall/infra/pb"
	"github.com/johnsmera/chall/infra/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 0, "Choose the server port")
	flag.Parse()

	log.Printf("start server on port %d", *port)

	userServer := setUpUserServer()

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userServer)

	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start listener server: ", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}

func setUpUserServer() *servers.UserServer {
	dbConnection := setUpDb()

	userRepository := repositories.UserRepositoryDb{Db: dbConnection}
	userServer := servers.NewUserServer()
	userServer.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}

	return userServer
}

func setUpDb() *gorm.DB {
	dbConnection := db.ConnectDB()
	dbConnection.LogMode(true)
	return dbConnection
}
