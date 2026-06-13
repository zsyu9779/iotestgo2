package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const DefaultToken = "valid-token-12345"

func UnaryInterceptor(token string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := ValidateIncomingContext(ctx, token); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func StreamInterceptor(token string) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := ValidateIncomingContext(stream.Context(), token); err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

func ValidateIncomingContext(ctx context.Context, token string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "missing metadata")
	}
	values := md.Get("authorization")
	if len(values) == 0 {
		return status.Error(codes.Unauthenticated, "missing authorization")
	}
	if !strings.HasPrefix(values[0], "Bearer ") {
		return status.Error(codes.Unauthenticated, "invalid authorization format")
	}
	got := strings.TrimPrefix(values[0], "Bearer ")
	if got != token {
		return status.Error(codes.PermissionDenied, "invalid token")
	}
	return nil
}
