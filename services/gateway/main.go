package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authpb "github.com/loak155/microservices/proto/auth"
)

func withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Run request", "http_method", r.Method, "http_url", r.URL)

		h.ServeHTTP(w, r)
	})
}

func newGateway(ctx context.Context) (http.Handler, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authEndpoint := fmt.Sprintf("%s:%s", os.Getenv("AUTH_SERVICE_HOST"), os.Getenv("AUTH_SERVICE_PORT"))
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, authEndpoint, opts)
	if err != nil {
		return nil, err
	}
	return mux, err
}

func main() {
	slog.Info("starting gateway")

	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux, err := newGateway(ctx)
	if err != nil {
		slog.Error("failed to create a new gateway", err)
	}

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		Handler: withLogger(mux),
	}
	go func() {
		defer s.Close()
		<-ctx.Done()
	}()

	s.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case v := <-quit:
		slog.Info("signal.Notify: ", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done: ", done)
	}
}
