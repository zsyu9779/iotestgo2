package auth

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestValidateIncomingContext(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer valid-token-12345"))
	if err := ValidateIncomingContext(ctx, DefaultToken); err != nil {
		t.Fatalf("expected valid token, got %v", err)
	}

	err := ValidateIncomingContext(context.Background(), DefaultToken)
	if status.Code(err) != codes.Unauthenticated {
		t.Fatalf("expected unauthenticated, got %v", status.Code(err))
	}

	bareTokenCtx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "valid-token-12345"))
	err = ValidateIncomingContext(bareTokenCtx, DefaultToken)
	if status.Code(err) != codes.Unauthenticated {
		t.Fatalf("expected unauthenticated for malformed bearer token, got %v", status.Code(err))
	}
}
