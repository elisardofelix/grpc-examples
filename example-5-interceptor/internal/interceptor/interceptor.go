package interceptor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/elisardofelix/grpc-examples/example-5-interceptor/internal/auth"
	"github.com/elisardofelix/grpc-examples/example-5-interceptor/proto"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	userIDKey = "user_id"
	claimsKey = "claims"
)

type (
	Validator interface {
		ValidateToken(ctx context.Context, token string) (string, error)
	}

	middleware struct {
		validator Validator
		secret    []byte
	}
)

func NewMiddleware(validator Validator, secret []byte) (*middleware, error) {
	if validator == nil {
		return nil, errors.New("validator cannot be nil")
	}

	return &middleware{validator: validator, secret: secret}, nil
}

func (m *middleware) UnaryAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// check the RPC method and only run if its the Protected RPC
	if info.FullMethod != proto.InterceptorService_Protected_FullMethodName {
		return handler(ctx, req)
	}

	// get the token from the metadata
	token, err := getTokenFromMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "token must be provided")
	}

	// call validate token
	userID, err := m.validator.ValidateToken(ctx, token)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidToken) {
			return nil, status.Error(codes.PermissionDenied, "invalid token")
		}

		slog.ErrorContext(ctx, "failed to validate token", "error", err)
		return nil, status.Error(codes.Internal, "error validating token")
	}

	// add the user ID to the context
	ctx = context.WithValue(ctx, userIDKey, userID)

	// call our handler
	return handler(ctx, req)
}

func (m *middleware) UnaryValidateMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// check the RPC method and only run if its the Validate RPC
	if info.FullMethod != proto.InterceptorService_Validate_FullMethodName {
		return handler(ctx, req)
	}

	token, err := getTokenFromMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "token must be provided")
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secret, nil
	})
	if err != nil {
		return "", status.Error(codes.PermissionDenied, "invalid token")
	}

	claimsMap := make(map[string]string)

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		for key, val := range claims {
			claim, ok := val.(string)
			if ok {
				claimsMap[key] = claim
				continue
			}

			ts, ok := val.(float64)
			if ok {
				claimsMap[key] = fmt.Sprintf("%.0f", ts)
			}
		}
	}
	ctx = context.WithValue(ctx, claimsKey, claimsMap)

	return handler(ctx, req)
}

func getTokenFromMetadata(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(meta["authorization"]) != 1 {
		return "", errors.New("token not found in metadata")
	}

	return meta["authorization"][0], nil
}
