package client

import (
	"context"
	"fmt"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CommonInterceptor struct{}

func NewCommonInterceptor() (*CommonInterceptor, error) {
	return &CommonInterceptor{}, nil
}

func (interceptor *CommonInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Info().Msg(fmt.Sprintf("---> unary interceptor %s", method))
		return invoker(interceptor.attachReqId(ctx), method, req, reply, cc, opts...)
	}
}

func (interceptor *CommonInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		log.Printf("--> stream interceptor: %s", method)
		return streamer(interceptor.attachReqId(ctx), desc, cc, method, opts...)
	}
}

func (interceptor *CommonInterceptor) attachReqId(ctx context.Context) context.Context {
	reqID := xid.New()
	log.Debug().Msgf("request-id: %s", reqID)
	return metadata.AppendToOutgoingContext(ctx, "Request-ID", reqID.String())
}
