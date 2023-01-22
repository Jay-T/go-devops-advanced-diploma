package server

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInteceptor struct {
	jwtManager       *JWTManager
	protectedMethods map[string]bool
}

func NewAuthInterceptor(jwtManager *JWTManager, protectedMethods map[string]bool) *AuthInteceptor {
	return &AuthInteceptor{jwtManager: jwtManager, protectedMethods: protectedMethods}
}

func (interceptor *AuthInteceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Info().Msg(fmt.Sprint("---> unary interceptor  ", info.FullMethod))

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInteceptor) authorize(ctx context.Context, method string) error {
	_, ok := interceptor.protectedMethods[method]

	if !ok {
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	_, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	log.Info().Msg(fmt.Sprintf("Request authorized for %s", method))
	return nil
}
