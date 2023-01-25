package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	"github.com/Jay-T/go-devops-advanced-diploma/internal/server"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := server.GetConfig()
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Error while getting config. %s", err.Error()))
	}

	if cfg.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	conn, err := sql.Open("postgres", cfg.DBAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewStore(conn)

	s, err := server.NewServer(ctx, cfg, store)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Could not get server. Error: %s", err))
	}

	go s.StartServer(ctx)

	log.Info().Msg(fmt.Sprintf("Listening socket: %s", cfg.Address))
	<-sigChan
	s.StopServer(ctx, cancel)
}
