package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/encryption"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/pbgen"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthUsecase interface {
	UserRegister(ctx context.Context, data *pbgen.RegisterUserRequest) (*pbgen.RegisterUserResponse, error)
	UserSignIn(ctx context.Context, email string) (*pbgen.UserSignInResponse, error)
	GetUserByCardID(ctx context.Context, cardID string) (*pbgen.GetUserByCardIDResponse, error)
}

type authUsecase struct {
	repo repo.AuthRepo
}

func NewAuthUsecase(repo repo.AuthRepo) authUsecase {
	return authUsecase{
		repo: repo,
	}
}

func checkUser(ctx context.Context, rp repo.AuthRepo, email string) (bool, error) {
	_, err := rp.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (uc authUsecase) UserRegister(ctx context.Context, data *pbgen.RegisterUserRequest) (*pbgen.RegisterUserResponse, error) {
	userExists, err := checkUser(ctx, uc.repo, data.GetEmail())
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, status.Error(codes.AlreadyExists, "user is AlreadyExists")
	}

	user, err := uc.repo.InsertUser(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pbgen.RegisterUserResponse{
		Msg:    "Success Insert User",
		Status: "Success",
		User: &pbgen.User{
			Id:                    int64(user.ID),
			CardId:                user.CardID,
			FullName:              user.FullName,
			Email:                 user.Email,
			Phone:                 user.Phone,
			Balance:               float64(user.Balance),
			Status:                user.Status,
			CurrentTripTerminalId: user.CurrentTripTerminalID,
			RegisteredAt:          timestamppb.New(user.RegisteredAt),
			UpdatedAt:             timestamppb.New(user.UpdatedAt),
		},
	}

	return res, nil
}

func (uc authUsecase) UserSignIn(ctx context.Context, email string) (*pbgen.UserSignInResponse, error) {
	user, err := uc.repo.GetUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(
			codes.NotFound,
			fmt.Sprintf("user with email %s is not found", email),
		)
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	token, err := encryption.GenerateToken(int64(user.ID), user.CardID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pbgen.UserSignInResponse{
		Msg:      "Success UserSignIn",
		Status:   "Success",
		JwtToken: &token,
	}

	return res, nil
}

func (uc authUsecase) GetUserByCardID(ctx context.Context, cardID string) (*pbgen.GetUserByCardIDResponse, error) {
	user, err := uc.repo.GetUserByCardID(ctx, cardID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(
			codes.NotFound,
			fmt.Sprintf("user with cardID %s is not found", cardID),
		)
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pbgen.GetUserByCardIDResponse{
		Msg:    "Success Insert User",
		Status: "Success",
		User: &pbgen.User{
			Id:                    int64(user.ID),
			CardId:                user.CardID,
			FullName:              user.FullName,
			Email:                 user.Email,
			Phone:                 user.Phone,
			Balance:               float64(user.Balance),
			Status:                user.Status,
			CurrentTripTerminalId: user.CurrentTripTerminalID,
			RegisteredAt:          timestamppb.New(user.RegisteredAt),
			UpdatedAt:             timestamppb.New(user.UpdatedAt),
		},
	}

	return res, nil
}
