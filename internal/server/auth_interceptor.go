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

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *serverStream) Context() context.Context {
	return s.ctx
}

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
		log.Debug().Msg(fmt.Sprint("---> unary interceptor  ", info.FullMethod))

		ctx, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInteceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Debug().Msg(fmt.Sprint("---> stream interceptor  ", info.FullMethod))

		ctx, err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, &serverStream{stream, ctx})
	}
}

func (interceptor *AuthInteceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	_, ok := interceptor.protectedMethods[method]

	if !ok {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	md.Append("username", claims.Username)
	ctx = metadata.NewIncomingContext(ctx, md)

	log.Info().Msgf("Request authorized for method: %s, user: %s", method, claims.Username)
	return ctx, nil
}
