package api

import (
	"context"
	"fmt"
	"github.com/jenson/county/core/models"
	"github.com/jenson/county/user/repository"
	"github.com/jenson/county/user/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	service.UnimplementedUserRpcServer
}

func (UserService) Login(ctx context.Context, in *models.LoginReq) (*models.LoginResp, error) {
	userRepo := repository.RepositoryInstance.UserRepository
	resp, err := userRepo.Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (UserService) SayHello(ctx context.Context, in *models.GreetReq) (*models.GreetResp, error) {

	return &models.GreetResp{
		Result: fmt.Sprintf("%s Hello", in.Username),
	}, nil
}

func (UserService) GetUserDetailByID(ctx context.Context, in *models.QueryUserDetailReq) (*models.QueryUserDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetailByID not implemented")
}
