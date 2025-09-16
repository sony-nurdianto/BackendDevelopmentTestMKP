package interceptor

import (
	"context"
	"fmt"
	"strings"

	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/encryption"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/pbgen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthServiceUnaryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	switch info.FullMethod {
	case pbgen.AuthService_GetUserByCardID_FullMethodName:
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println("No metadata found")
			break
		}

		authHeader := ""
		if vals, exists := md["authorization"]; exists && len(vals) > 0 {
			authHeader = vals[0]
		}

		if authHeader == "" {
			fmt.Println("Warning: Authorization header kosong!")
			return nil, status.Error(codes.Unauthenticated, "Authorization is empty ,Unauthenticated")
		}

		authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println(authHeader)

		_, err := encryption.VerifyToken(
			authHeader,
		)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
	}

	return handler(ctx, req)
}
