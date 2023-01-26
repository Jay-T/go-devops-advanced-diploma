package client

import (
	"context"
	"time"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service  pb.AuthenticationClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthenticationClient(cc)
	return &AuthClient{service, username, password}
}

func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Login:    client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetToken(), nil
}
