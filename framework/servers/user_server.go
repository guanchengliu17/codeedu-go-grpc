package servers

import (
	"context"
	"log"

	"github.com/leopedroso45/codeedu/application/usecases"
	"github.com/leopedroso45/codeedu/domain"
	"github.com/leopedroso45/codeedu/framework/pb"
)

type UserServer struct {
	User        domain.User
	UserUseCase usecases.UserUseCase
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (UserServer *UserServer) CrateUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	UserServer.User.Name = request.GetName()
	UserServer.User.Email = request.GetEmail()
	UserServer.User.Password = request.GetPassword()
	userResp, err := UserServer.UserUseCase.Create(&UserServer.User)
	if err != nil {
		log.Fatalf("Error during RPC Create User:  %v", err)
	}

	return &pb.UserResponse{
		Token: userResp.Token,
	}, nil

}
