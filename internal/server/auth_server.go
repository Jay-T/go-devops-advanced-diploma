package server

import (
	"context"
	"fmt"

	pb "github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthenticationServer
	userStore  UserStore
	jwtManager *JWTManager
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{pb.UnimplementedAuthenticationServer{}, userStore, jwtManager}
}

func (s *AuthServer) UserSignIn(ctx context.Context, in *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	log.Info().Msg(fmt.Sprintf("Got SignIn request for login '%s', password '%s'", in.Login, in.Password))

	user, err := s.userStore.Find(in.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(in.Password) {
		return nil, status.Error(codes.NotFound, "username/password incorrect")
	}

	token, err := s.jwtManager.GeneratetToken(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.UserSignInResponse{Token: token}
	return res, nil
}

func (s *AuthServer) UserSignUp(ctx context.Context, in *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	log.Info().Msg(fmt.Sprintf("Got SignUp request for login '%s', password '%s'", in.Login, in.Password))

	// u := UserFromSignUpPB(in)
	// err := u.registerNewUser(ctx)
	// if err != nil {
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	// token, err := u.authenticateUser(ctx, s.tokenLifeTime)
	// if err != nil {
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }
	user, err := s.userStore.Find(in.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(in.Password) {
		return nil, status.Error(codes.NotFound, "username/password incorrect")
	}

	token, err := s.jwtManager.GeneratetToken(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &pb.UserSignUpResponse{
		Token: token,
	}, nil
}
