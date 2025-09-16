package service

import (
	"context"

	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/pbgen"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/usecase"
)

type AuthServiceServer struct {
	pbgen.UnimplementedAuthServiceServer
	usecase usecase.AuthUsecase
}

func NewAuthServiceServer(uc usecase.AuthUsecase) *AuthServiceServer {
	return &AuthServiceServer{
		usecase: uc,
	}
}

func (s AuthServiceServer) RegisterUser(ctx context.Context, in *pbgen.RegisterUserRequest) (*pbgen.RegisterUserResponse, error) {
	return s.usecase.UserRegister(ctx, in)
}

func (s AuthServiceServer) UserSignIn(ctx context.Context, in *pbgen.UserSignInRequest) (*pbgen.UserSignInResponse, error) {
	return s.usecase.UserSignIn(ctx, in.GetEmail())
}

func (s AuthServiceServer) GetUserByCardID(ctx context.Context, in *pbgen.GetUserByCardIDRequest) (*pbgen.GetUserByCardIDResponse, error) {
	return s.usecase.GetUserByCardID(ctx, in.GetCardId())
}

func (s AuthServiceServer) Greeter(ctx context.Context, in *pbgen.GreeterRequest) (*pbgen.GreeterResponse, error) {
	return &pbgen.GreeterResponse{
		Hallo: in.GetName(),
	}, nil
}
