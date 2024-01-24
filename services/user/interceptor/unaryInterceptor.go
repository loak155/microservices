package interceptor

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
)

func LoggingInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	slog.Info("Method: ", info.FullMethod, ", Request: ", req)
	res, err := handler(ctx, req)
	slog.Info("Response: ", res, "Error: ", err)
	return res, err
}
