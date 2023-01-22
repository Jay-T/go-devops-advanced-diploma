package server

import (
	"context"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AnythingElseServer struct {
	pb.UnimplementedAnythingElseServer
}

func NewAnythingElseServer() *AnythingElseServer {
	return &AnythingElseServer{pb.UnimplementedAnythingElseServer{}}
}

func (s *AnythingElseServer) GetUserInfo(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	log.Info().Msg("received request for AnythingElseServer.GetUserInfo")

	return &emptypb.Empty{}, nil
}
