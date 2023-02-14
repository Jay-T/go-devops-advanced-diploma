package client

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Work with files.",
	Long:  `Work with files.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("running file!")
	},
}

func init() {
	log.Info().Msgf("Calling file")
	rootCmd.AddCommand(fileCmd)
}
