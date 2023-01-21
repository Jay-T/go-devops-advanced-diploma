package server

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ Server = (*GRPCServer)(nil)

type User struct {
	Username string `json:"login"`
	Password string `json:"password"`
	Passhash string
}

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
	pb.UnimplementedGophKeeperServer
}

func NewServer(ctx context.Context, cfg *Config) (Server, error) {
	genericService, err := NewGenericServer(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &GRPCServer{
		genericService,
		pb.UnimplementedGophKeeperServer{},
	}, nil
}

func (s *GRPCServer) StartServer(ctx context.Context) {
	listen, err := net.Listen("tcp", s.Cfg.Address)
	if err != nil {
		log.Fatal().Err(err).Str("func", "StartServer")
	}

	// interceptors := []grpc.UnaryServerInterceptor{
	// 	s.checkReqIDInterceptor,
	// }
	interceptors := []grpc.UnaryServerInterceptor{}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors...))
	pb.RegisterGophKeeperServer(server, s)
	reflection.Register(server)

	go func() {
		log.Info().Msg(fmt.Sprintf("Starting GRPC server on socket %s", s.Cfg.Address))
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

func (s *GRPCServer) UserSignIn(ctx context.Context, in *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	log.Info().Msg(fmt.Sprintf("Got SignIn request for login '%s', password '%s'", in.Login, in.Password))

	return &pb.UserSignInResponse{
		Token: "123456",
	}, nil
}

func (s *GRPCServer) UserSignUp(ctx context.Context, in *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	log.Info().Msg(fmt.Sprintf("Got SignUp request for login '%s', password '%s'", in.Login, in.Password))

	return &pb.UserSignUpResponse{
		Token: "123456",
	}, nil
}
