package servers

import (
	"context"
	"log"

	"github.com/johnsmera/chall/application/repositories/usecases"
	"github.com/johnsmera/chall/domain"
	"github.com/johnsmera/chall/infra/pb"
)

type UserServer struct {
	User        domain.User
	UserUseCase usecases.UserUseCase
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (UserServer *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	UserServer.User.Name = req.GetName()
	UserServer.User.Email = req.GetEmail()
	UserServer.User.Password = req.GetPassword()

	user, err := UserServer.UserUseCase.Create(&UserServer.User)

	if err != nil {
		log.Fatalf("Error during user rpc create %v", err)
	}

	return &pb.UserResponse{
		Token: user.Token,
	}, nil
}
