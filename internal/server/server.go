package server

import (
	"context"
	"fmt"
	"net"
	"os"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	pb "github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	secretKey = "secret_key"
)

var _ Server = (*GRPCServer)(nil)

type GenericService struct {
	Cfg *Config
	// Encryptor     *Eecryptor
	// Decryptor     *Decryptor
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
	pb.UnimplementedAnythingElseServer
}

func NewServer(ctx context.Context, cfg *Config, store db.Store) (Server, error) {
	genericService, err := NewGenericServer(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &GRPCServer{
		genericService,
		store,
		pb.UnimplementedAnythingElseServer{},
	}, nil
}
func protectedMethods() map[string]bool {
	const protectedServicePath = "/go_devops_advanced_diploma.AnythingElse/"
	return map[string]bool{
		protectedServicePath + "GetUserInfo": true,
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

	anythingElseServer := NewAnythingElseServer()

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))

	pb.RegisterAnythingElseServer(server, anythingElseServer)
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
