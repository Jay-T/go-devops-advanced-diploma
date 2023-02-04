package server

import (
	"context"
	"database/sql"
	"encoding/base32"
	"encoding/base64"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	"github.com/Jay-T/go-devops-advanced-diploma/internal/crypto"
	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// getUsernameFromContext extracts username from incoming metadata.
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

// SecretServer struct
type SecretServer struct {
	secretStore db.Store
	crypto      *crypto.CryptoService
	pb.UnimplementedSecretServer
}

// NewSecretServer returns new SecretServer instance.
func NewSecretServer(secretStore db.Store, CS *crypto.CryptoService) *SecretServer {
	return &SecretServer{
		secretStore,
		CS,
		pb.UnimplementedSecretServer{},
	}
}

// CreateSecret creates secret in secret storage.
func (s *SecretServer) CreateSecret(ctx context.Context, in *pb.CreateSecretRequest) (*pb.CreateSecretResponse, error) {
	account, err := findAccount(ctx, s.secretStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("Got CreateSecret request for login '%s'", account.Username)

	masterKey := base32.StdEncoding.EncodeToString([]byte(in.Data.Masterkey))

	cipher, err := s.crypto.Encrypt(in.Data.Value, []byte(masterKey))
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot encrypt the secret. Err :%s", err))
	}

	arg := db.CreateSecretParams{
		AccountID: account.ID,
		Key:       in.Data.Key,
		Value:     cipher,
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
		Value: string(secret.Value),
	}

	return &pb.CreateSecretResponse{
		Data: message,
	}, nil
}

// DeleteSecret deletes secret from secret storage.
func (s *SecretServer) DeleteSecret(ctx context.Context, in *pb.DeleteSecretRequest) (*pb.DeleteSecretResponse, error) {
	account, err := findAccount(ctx, s.secretStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("Got DeleteSecret request for login '%s'", account.Username)

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

// GetSecret returns secret info from secret storage.
func (s *SecretServer) GetSecret(ctx context.Context, in *pb.GetSecretRequest) (*pb.GetSecretResponse, error) {
	account, err := findAccount(ctx, s.secretStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("Got GetSecret request for login '%s'", account.Username)

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

	masterKey := base32.StdEncoding.EncodeToString([]byte(in.Masterkey))

	decipher, err := s.crypto.Decrypt(secret.Value, []byte(masterKey))
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot decrypt the secret. Err: %s", err))
	}

	secretMSG := &pb.SecretMessage{
		Key:   secret.Key,
		Value: decipher,
	}

	return &pb.GetSecretResponse{
		Data: secretMSG,
	}, nil
}

// ListSecret returns all secrets from secret storage for user.
func (s *SecretServer) ListSecret(ctx context.Context, in *pb.ListSecretRequest) (*pb.ListSecretResponse, error) {
	account, err := findAccount(ctx, s.secretStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err: %s", err))
	}

	log.Info().Msgf("Got ListSecret request for login '%s'", account.Username)

	secrets, err := s.secretStore.ListSecrets(ctx, account.ID)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get secrets from db. Err: %s", err))
	}

	secretsList := []*pb.SecretMessage{}
	masterKey := base32.StdEncoding.EncodeToString([]byte(in.Masterkey))
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot encode masterkey. Err:%s", err))
	}

	for _, i := range secrets {
		decipher, err := s.crypto.Decrypt(i.Value, []byte(masterKey))
		if err != nil {
			if err.Error() == "cipher: message authentication failed" {
				return nil, logError(status.Error(codes.InvalidArgument, "masterkey is not correct."))
			}
			return nil, logError(status.Errorf(codes.Internal, "cannot decrypt the secret. Err: %s", err))
		}

		secretsList = append(secretsList, &pb.SecretMessage{
			Key:   i.Key,
			Value: decipher,
		})
	}

	return &pb.ListSecretResponse{
		Data: secretsList,
	}, nil
}

// UpdateSecret updates secret info in secret storage.
func (s *SecretServer) UpdateSecret(ctx context.Context, in *pb.UpdateSecretRequest) (*pb.UpdateSecretResponse, error) {
	account, err := findAccount(ctx, s.secretStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("Got UpdateSecret request for login '%s'", account.Username)

	masterKey, err := base64.StdEncoding.DecodeString(in.Data.Masterkey)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot decode masterkey. Err :%s", err))
	}

	cipher, err := s.crypto.Encrypt(in.Data.Value, masterKey)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot encrypt the secret. Err :%s", err))
	}

	arg := db.UpdateSecretParams{
		Key:       in.Data.Key,
		AccountID: account.ID,
		Value:     cipher,
	}

	if err := s.secretStore.UpdateSecret(ctx, arg); err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot update the secret. Err :%s", err))
	}
	return &pb.UpdateSecretResponse{
		Data: &pb.SecretMessage{
			Key:   in.Data.Key,
			Value: in.Data.Value,
		},
	}, nil
}
