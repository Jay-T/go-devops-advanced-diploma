package server

import (
	"context"
	"fmt"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	pb "github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/Jay-T/go-devops-advanced-diploma/internal/util"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// findAccount gets account info from Accounts table.
func findAccount(ctx context.Context, s db.Store) (db.Account, error) {
	username, err := getUsernameFromContext(ctx)
	if err != nil {
		return db.Account{}, err
	}

	account, err := s.GetAccount(ctx, username)
	if err != nil {
		return db.Account{}, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	return account, nil
}

type AuthServer struct {
	pb.UnimplementedAuthenticationServer
	accountStore db.Store
	jwtManager   *JWTManager
}

func NewAuthServer(store db.Store, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{pb.UnimplementedAuthenticationServer{}, store, jwtManager}
}

func (s *AuthServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Info().Msg(fmt.Sprintf("Got SignIn request for login '%s', password '%s'", in.Login, in.Password))

	acc, err := s.accountStore.GetAccount(ctx, in.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if !acc.IsCorrectPassword(in.Password) {
		return nil, status.Error(codes.NotFound, "username/password incorrect")
	}

	token, expirationTime, err := s.jwtManager.GeneratetToken(&acc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &pb.LoginResponse{
		Login:                acc.Username,
		Token:                token,
		AccessTokenExpiresAt: timestamppb.New(expirationTime),
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Info().Msg(fmt.Sprintf("Got SignUp request for login '%s', password '%s'", in.Login, in.Password))

	hash, err := util.HashPassword(in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password: %v", err)
	}

	arg := db.CreateAccountParams{
		Username: in.Login,
		Passhash: hash,
	}

	newAcc, err := s.accountStore.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	token, expirationTime, err := s.jwtManager.GeneratetToken(&newAcc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &pb.RegisterResponse{
		Login:                newAcc.Username,
		Token:                token,
		AccessTokenExpiresAt: timestamppb.New(expirationTime),
	}, nil
}
