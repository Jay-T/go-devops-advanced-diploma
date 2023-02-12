package client

import (
	"io/ioutil"
	"os"
	"path/filepath"

	pb "github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCAgent struct {
	authClient   *AuthClient
	secretClient pb.SecretClient
	fileClient   pb.FileClient
	token        string
	Config       *Config
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

// NewGRPCAgent returns GRPCAgent for work.
func NewGRPCAgent(cfg *Config) (*GRPCAgent, error) {

	c := &GRPCAgent{}
	c.Config = cfg
	cc1, err := grpc.Dial(c.Config.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	authClient := NewAuthClient(cc1, c.Config.Tokenfile)
	authInterceptor, err := NewAuthInterceptor(authClient, protectedMethods())
	if err != nil {
		log.Err(err).Msgf("error during authInterceptor initialization.")
		return nil, err
	}

	commonInterceptor, err := NewCommonInterceptor()
	if err != nil {
		log.Err(err).Msgf("error during commonInterceptor initialization.")
		return nil, err
	}

	cc2, err := grpc.Dial(
		c.Config.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(authInterceptor.Unary(), commonInterceptor.Unary()),
		grpc.WithChainStreamInterceptor(authInterceptor.Stream(), commonInterceptor.Stream()),
	)

	if err != nil {
		log.Fatal().Msgf("cannot dial server", err)
	}

	c.authClient = authClient
	c.fileClient = pb.NewFileClient(cc2)
	c.secretClient = pb.NewSecretClient(cc2)

	return c, nil
}

func (ga *GRPCAgent) SaveToken(token string) error {
	ga.token = token
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(home, ga.Config.Tokenfile), []byte(token), 0600)
}
