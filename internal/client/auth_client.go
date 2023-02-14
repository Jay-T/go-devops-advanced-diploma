package client

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service pb.AuthenticationClient
	token   string
}

func NewAuthClient(cc *grpc.ClientConn, tokenfile string) *AuthClient {
	ac := &AuthClient{}
	ac.service = pb.NewAuthenticationClient(cc)
	ac.LoadToken(tokenfile)
	return ac
}

func (ac *AuthClient) LoadToken(tokenfile string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	tokenBytes, err := ioutil.ReadFile(filepath.Join(home, tokenfile))
	if err != nil {
		return err
	}

	ac.token = string(tokenBytes)
	return nil
}
