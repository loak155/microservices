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

	"github.com/loak155/microservices/services/article/db"
	"github.com/loak155/microservices/services/article/repository"
	"github.com/loak155/microservices/services/article/router"
	"github.com/loak155/microservices/services/article/usecase"
	"github.com/loak155/microservices/services/article/validator"

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
	articleValidator := validator.NewArticleValidator()
	articleRepository := repository.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepository, articleValidator)
	router.NewArticleGRPCServer(server, articleUsecase)

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
