package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/pbgen"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/repo"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/service"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	authRepo, err := repo.NewAuthRepo(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	authUsecase := usecase.NewAuthUsecase(authRepo)

	svc := service.NewAuthServiceServer(authUsecase)
	svr := grpc.NewServer()

	pbgen.RegisterAuthServiceServer(svr, svc)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	go func(lis net.Listener) {
		if err := svr.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}(listener)

	var once sync.Once

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Server Stoping, Gracefully Stop...")
			svr.GracefulStop()
			fmt.Println("Aplication Quit")
			return
		default:
			once.Do(func() { fmt.Println("Server Is Running At 0.0.0.0:50051") })

		}
	}
}
