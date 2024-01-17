package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/loak155/microservices/services/user/db"
	"github.com/loak155/microservices/services/user/repository"
	"github.com/loak155/microservices/services/user/router"
	"github.com/loak155/microservices/services/user/usecase"
	"github.com/loak155/microservices/services/user/validator"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	slog.Info("starting grpc server")

	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := grpc.NewServer()
	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	router.NewUserGRPCServer(server, userUsecase)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	if err != nil {
		slog.Error("failed to listen to address")
		cancel()
	}
	err = server.Serve(listener)
	if err != nil {
		slog.Error("failed to start gRPC server")
		cancel()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case v := <-quit:
		slog.Info("signal.Notify: ", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done: ", done)
	}
}
