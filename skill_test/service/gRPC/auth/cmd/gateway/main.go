package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/pbgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	mux := runtime.NewServeMux()
	err := pbgen.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:50051",
		opts,
	)
	if err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalln(err)
	}
}
