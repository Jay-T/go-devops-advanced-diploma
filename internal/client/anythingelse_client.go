package client

import (
	"context"
	"time"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AnythingElseClient struct {
	service pb.AnythingElseClient
}

func NewAnythingElseClient(cc *grpc.ClientConn) *AnythingElseClient {
	service := pb.NewAnythingElseClient(cc)
	return &AnythingElseClient{service}
}

func (c *AnythingElseClient) GetUserInfo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info().Msg("calling AnythingElseClient.GetUserInfo")

	req := &emptypb.Empty{}
	_, err := c.service.GetUserInfo(ctx, req)

	if err != nil {
		return err
	}

	log.Info().Msg("got reply from AnythingElseClient.GetUserInfo")
	return nil
}
