package server

import (
	"context"
	"database/sql"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func getUsernameFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", logError(status.Errorf(codes.Unauthenticated, "metadata is not provided"))
	}

	values := md["username"]
	if len(values) == 0 {
		return "", logError(status.Errorf(codes.Internal, "no username found in context"))
	}

	return values[0], nil
}

type SecretServer struct {
	secretStore db.Store
	pb.UnimplementedSecretServer
}

func NewSecretServer(secretStore db.Store) *SecretServer {
	return &SecretServer{secretStore, pb.UnimplementedSecretServer{}}
}

func (s *SecretServer) CreateSecret(ctx context.Context, in *pb.CreateSecretRequest) (*pb.CreateSecretResponse, error) {
	username, err := getUsernameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("Got CreateSecret request for login '%s'", username)

	account, err := s.secretStore.GetAccount(ctx, username)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	// TODO(): ADD VALUE ENCRYPTION HERE!

	arg := db.CreateSecretParams{
		AccountID: account.ID,
		Key:       in.Data.Key,
		Value:     in.Data.Value,
	}

	secret, err := s.secretStore.CreateSecret(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, logError(status.Errorf(codes.AlreadyExists, "Secret already exists: %s", err))
			}
		}
		return nil, logError(status.Errorf(codes.Internal, "failed to create secret: Err: %s", err))
	}

	message := &pb.SecretMessage{
		Key:   secret.Key,
		Value: secret.Value,
	}

	return &pb.CreateSecretResponse{
		Data: message,
	}, nil
}

func (s *SecretServer) UpdateSecret(context.Context, *pb.UpdateSecretRequest) (*pb.UpdateSecretResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Uninmplemented")
}

func (s *SecretServer) DeleteSecret(ctx context.Context, in *pb.DeleteSecretRequest) (*pb.DeleteSecretResponse, error) {
	username, err := getUsernameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	account, err := s.secretStore.GetAccount(ctx, username)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("Got DeleteSecret request for login '%s'", username)

	arg := db.GetSecretParams{
		Key:       in.Key,
		AccountID: account.ID,
	}
	secret, err := s.secretStore.GetSecret(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, logError(status.Error(codes.NotFound, "cannot find secret"))
		}

		return nil, logError(status.Errorf(codes.Internal, "cannot get secret: Err: %s", err))
	}

	arg2 := db.DeleteSecretParams{
		Key:       secret.Key,
		AccountID: account.ID,
	}

	err = s.secretStore.DeleteSecret(ctx, arg2)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot delete secret db. Err: %s", err))
	}

	return &pb.DeleteSecretResponse{
		Key: secret.Key,
	}, nil
}

func (s *SecretServer) GetSecret(context.Context, *pb.GetSecretRequest) (*pb.GetSecretResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Uninmplemented")
}

func (s *SecretServer) ListSecret(context.Context, *pb.ListSecretRequest) (*pb.ListSecretResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Uninmplemented")
}
