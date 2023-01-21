package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := server.GetConfig()
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Error while getting config. %s", err.Error()))
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	s, err := server.NewServer(ctx, cfg)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Could not run GRPC server. Error: %s", err))
	}

	go s.StartServer(ctx)

	log.Info().Msg(fmt.Sprintf("Listening socket: %s", cfg.Address))
	<-sigChan
	s.StopServer(ctx, cancel)
}
