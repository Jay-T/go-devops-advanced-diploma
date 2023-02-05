package server

import (
	"context"
	"fmt"
	"net"
	"os"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	"github.com/Jay-T/go-devops-advanced-diploma/internal/crypto"
	pb "github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	secretKey = "secret_key"
	cypherKey = "VeryStrongKey"
)

var _ Server = (*GRPCServer)(nil)

type GenericService struct {
	Cfg *Config
	// backuper      StorageBackuper
}

func NewGenericServer(ctx context.Context, cfg *Config) (*GenericService, error) {
	var s GenericService

	s.Cfg = cfg
	return &s, nil
}

type Server interface {
	StartServer(ctx context.Context)
	StopServer(ctx context.Context, cancel context.CancelFunc)
}
type GRPCServer struct {
	*GenericService
	store db.Store
}

func NewServer(ctx context.Context, cfg *Config, store db.Store) (Server, error) {
	genericService, err := NewGenericServer(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &GRPCServer{
		genericService,
		store,
	}, nil
}
func protectedMethods() map[string]bool {
	const (
		protectedSecretServicePath = "/go_devops_advanced_diploma.Secret/"
		protectedFileServicePath   = "/go_devops_advanced_diploma.File/"
	)
	return map[string]bool{
		protectedSecretServicePath + "CreateSecret": true,
		protectedSecretServicePath + "DeleteSecret": true,
		protectedSecretServicePath + "GetSecret":    true,
		protectedSecretServicePath + "ListSecret":   true,
		protectedSecretServicePath + "UpdateSecret": true,
		protectedFileServicePath + "CreateFile":     true,
		protectedFileServicePath + "DeleteFile":     true,
		protectedFileServicePath + "GetFileInfo":    true,
		protectedFileServicePath + "DownloadFile":   true,
		protectedFileServicePath + "ListFiles":      true,
		protectedFileServicePath + "UpdateFileName": true,
	}
}

func (s *GRPCServer) StartServer(ctx context.Context) {
	listen, err := net.Listen("tcp", s.Cfg.Address)
	if err != nil {
		log.Fatal().Err(err).Str("func", "StartServer")
	}

	// accountStore := NewInMemoryAccountStore()
	// err = seedAccounts(accountStore)
	if err != nil {
		log.Fatal().Msg("cannot seed users")
	}
	jwtManager := NewJWTManager(secretKey, s.Cfg.TokenLifeTime)
	authServer := NewAuthServer(s.store, jwtManager)
	interceptor := NewAuthInterceptor(jwtManager, protectedMethods())
	cryptoService := crypto.NewCryptoService(cypherKey)

	secretServer := NewSecretServer(s.store, cryptoService)
	fileServer := NewFileServer(ctx, s.store)

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	server := grpc.NewServer(serverOptions...)

	pb.RegisterSecretServer(server, secretServer)
	pb.RegisterFileServer(server, fileServer)
	pb.RegisterAuthenticationServer(server, authServer)
	reflection.Register(server)

	go func() {
		log.Info().Msg(fmt.Sprintf("Starting GRPC server with following config: %+v", s.Cfg))
		if err := server.Serve(listen); err != nil {
			log.Fatal().Err(err).Str("func", "StartServer")
		}
	}()

	<-ctx.Done()
	log.Info().Msg("Finished to serve gRPC requests")
}

func (s *GRPCServer) StopServer(ctx context.Context, cancel context.CancelFunc) {
	log.Info().Msg("Received a SIGINT! Stopping application")
	cancel()
	log.Info().Msg("Canceled all goroutines.")
	os.Exit(1)
}
