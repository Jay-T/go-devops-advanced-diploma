package main

import (
	"time"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/client"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	username        = "admin"
	password        = "admin"
	refreshDuration = 15 * time.Second
)

func authMethods() map[string]bool {
	const protectedServicePath = "/go_devops_advanced_diploma.AnythingElse/"
	return map[string]bool{
		protectedServicePath + "GetUserInfo": true,
	}
}

func testClient(c *client.AnythingElseClient, duration time.Duration) {
	wait := duration
	for {
		time.Sleep(wait)
		c.GetUserInfo()
	}
}

func main() {
	serverAddress := "localhost:53000"

	// cc1, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	cc1, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msg("cannot not dial to the server.")
	}

	authClient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal().Msg("cannot create auth interceptor.")
	}

	cc2, err := grpc.Dial(
		serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		log.Fatal().Msg("cannot not dial to the server.")
	}

	anythingElseClient := client.NewAnythingElseClient(cc2)
	testClient(anythingElseClient, time.Second*3)
}
