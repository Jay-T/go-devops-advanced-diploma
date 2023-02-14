package client

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CtxKey string

var (
	runtimeViper = viper.New()
	cfgFile      string
	rootCmd      = &cobra.Command{
		Use:   "gophclient",
		Short: "Client to work with gophkeeper service.",
		Long:  `Client to work with gophkeeper service.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Debug().Msgf("running %s", cmd.Name())
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		runtimeViper.SetConfigFile(cfgFile)
	} else {
		runtimeViper.AddConfigPath(".")
		runtimeViper.AddConfigPath("./cmd/client/.")
		runtimeViper.SetConfigType("yaml")
		runtimeViper.SetConfigName("config")
	}

	runtimeViper.AutomaticEnv()

	err := runtimeViper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", runtimeViper.ConfigFileUsed())
	}

	cfg := &Config{}
	err = runtimeViper.Unmarshal(cfg)
	if err != nil {
		fmt.Println(err)
	}

	rootCtx := context.Background()

	c, err := NewGRPCAgent(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("error trying to initialize GRPCAgent.")
	}

	ctx := context.WithValue(rootCtx, CtxKey("c"), c)
	rootCmd.SetContext(ctx)

}
