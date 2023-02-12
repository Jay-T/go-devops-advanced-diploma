package client

import (
	"io/ioutil"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service pb.AuthenticationClient
	token   string
}

func NewAuthClient(cc *grpc.ClientConn) *AuthClient {
	ac := &AuthClient{}
	ac.service = pb.NewAuthenticationClient(cc)
	ac.LoadToken()
	return ac
}

func (ac *AuthClient) LoadToken() error {
	tokenBytes, err := ioutil.ReadFile(".token")
	if err != nil {
		return err
	}

	ac.token = string(tokenBytes)
	return nil
}
