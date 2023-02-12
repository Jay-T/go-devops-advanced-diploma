package client

import (
	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/prototext"
)

var secretCmd = &cobra.Command{
	Use:              "secret",
	Short:            "Work with secrets.",
	Long:             `Work with secrets.`,
	TraverseChildren: true,
}

var secretListCmd = &cobra.Command{
	Use:   "list [-m <MASTERKEY>]",
	Short: "List all secrets.",
	Long:  `List all secrets.`,
	RunE:  SecretList,
}

var secretCreateCmd = &cobra.Command{
	Use:   "create [-m <MASTERKEY>] -k <KEY> -v <VALUE> [--metadata  KEY=VALUE[,KEY=VALUE...]]",
	Short: "Create a secret.",
	Long:  `Create a secret..`,
	RunE:  SecretCreate,
}

var secretDeleteCmd = &cobra.Command{
	Use:   "delete -k <KEY>",
	Short: "Delete a secret.",
	Long:  `Delete a secret.`,
	RunE:  SecretDelete,
}

var secretGetCmd = &cobra.Command{
	Use:   "get [-m <MASTERKEY>] -k <KEY>",
	Short: "Get a secret.",
	Long:  `Get a secret.`,
	RunE:  SecretGet,
}

var secretUpdateCmd = &cobra.Command{
	Use:   "update [-m <MASTERKEY>] -k <KEY> -v <VALUE> [--metadata  KEY=VALUE[,KEY=VALUE...]]",
	Short: "Update a secret.",
	Long:  `Update a secret.`,
	RunE:  SecretUpdate,
}

func SecretList(cmd *cobra.Command, args []string) error {
	log.Debug().Msgf("running %s", cmd.Name())
	ctx := cmd.Root().Context()
	c := ctx.Value(CtxKey("c")).(*GRPCAgent)

	req := &pb.ListSecretRequest{
		Masterkey: c.Config.Masterkey,
	}

	resp, err := c.secretClient.ListSecret(ctx, req)
	if err != nil {
		log.Err(err).Msg("cannot get secrets from server.")
	}

	log.Info().Msgf("Response: \n%s", prototext.Format(resp))
	return nil
}

func SecretCreate(cmd *cobra.Command, args []string) error {
	log.Debug().Msgf("running %s", cmd.Name())
	k, err := cmd.Flags().GetString("key")
	if err != nil {
		return err
	}
	v, err := cmd.Flags().GetString("value")
	if err != nil {
		return err
	}
	md, _ := cmd.Flags().GetStringToString("metadata")
	if err != nil {
		return err
	}

	ctx := cmd.Root().Context()
	c := ctx.Value(CtxKey("c")).(*GRPCAgent)

	req := &pb.CreateSecretRequest{
		Data: &pb.SecretMessage{
			Key:       k,
			Value:     v,
			Masterkey: c.Config.Masterkey,
		},
	}
	if len(md) > 0 {
		for mdk, mdv := range md {
			req.Data.Metadata = append(req.Data.Metadata, &pb.Metadata{
				Key:   mdk,
				Value: mdv,
			})
		}
	}

	resp, err := c.secretClient.CreateSecret(ctx, req)
	if err != nil {
		log.Err(err).Msg("cannot create secret.")
		return err
	}

	log.Info().Msgf("successfully created secret. \n%s", prototext.Format(resp))
	return nil
}

func SecretDelete(cmd *cobra.Command, args []string) error {
	log.Debug().Msgf("running %s", cmd.Name())
	k, err := cmd.Flags().GetString("key")
	if err != nil {
		return err
	}
	ctx := cmd.Root().Context()
	c := ctx.Value(CtxKey("c")).(*GRPCAgent)

	req := &pb.DeleteSecretRequest{
		Key: k,
	}

	resp, err := c.secretClient.DeleteSecret(ctx, req)
	if err != nil {
		log.Err(err).Msg("cannot delete secret.")
		return err
	}

	log.Info().Msgf("successfully deleted secret. \n%s", prototext.Format(resp))
	return nil
}

func SecretGet(cmd *cobra.Command, args []string) error {
	log.Debug().Msgf("running %s", cmd.Name())
	k, err := cmd.Flags().GetString("key")
	if err != nil {
		return err
	}
	ctx := cmd.Root().Context()
	c := ctx.Value(CtxKey("c")).(*GRPCAgent)

	req := &pb.GetSecretRequest{
		Key:       k,
		Masterkey: c.Config.Masterkey,
	}

	resp, err := c.secretClient.GetSecret(ctx, req)
	if err != nil {
		log.Err(err).Msg("cannot get secret.")
		return err
	}

	log.Info().Msgf("successfully got secret. \n%s", prototext.Format(resp))
	return nil
}

func SecretUpdate(cmd *cobra.Command, args []string) error {
	log.Debug().Msgf("running %s", cmd.Name())
	k, err := cmd.Flags().GetString("key")
	if err != nil {
		return err
	}
	v, err := cmd.Flags().GetString("value")
	if err != nil {
		return err
	}
	md, _ := cmd.Flags().GetStringToString("metadata")
	if err != nil {
		return err
	}

	ctx := cmd.Root().Context()
	c := ctx.Value(CtxKey("c")).(*GRPCAgent)

	req := &pb.UpdateSecretRequest{
		Data: &pb.SecretMessage{
			Key:       k,
			Value:     v,
			Masterkey: c.Config.Masterkey,
		},
	}
	if len(md) > 0 {
		for mdk, mdv := range md {
			req.Data.Metadata = append(req.Data.Metadata, &pb.Metadata{
				Key:   mdk,
				Value: mdv,
			})
		}
	}

	resp, err := c.secretClient.UpdateSecret(ctx, req)
	if err != nil {
		log.Err(err).Msg("cannot update secret.")
		return err
	}

	log.Info().Msgf("successfully created update. \n%s", prototext.Format(resp))
	return nil
}

func init() {
	var m, k, v string
	var md map[string]string

	secretCmd.PersistentFlags().StringVarP(&m, "masterkey", "m", "", "Masterkey for secrets encryption on server.")
	runtimeViper.BindPFlag("masterkey", secretCmd.Flag("masterkey"))

	secretCreateCmd.Flags().StringVarP(&k, "key", "k", "", "Secret name")
	secretCreateCmd.Flags().StringVarP(&v, "value", "v", "", "Secret value")
	secretCreateCmd.Flags().StringToStringVarP(&md, "metadata", "", nil, "Secret metadata key=value")
	secretCreateCmd.MarkFlagsRequiredTogether("key", "value")

	secretDeleteCmd.Flags().StringVarP(&k, "key", "k", "", "Secret name")

	secretGetCmd.Flags().StringVarP(&k, "key", "k", "", "Secret name")

	secretUpdateCmd.Flags().StringVarP(&k, "key", "k", "", "Secret name")
	secretUpdateCmd.Flags().StringVarP(&v, "value", "v", "", "Secret value")
	secretUpdateCmd.Flags().StringToStringVarP(&md, "metadata", "", nil, "Secret metadata key=value")
	secretUpdateCmd.MarkFlagsRequiredTogether("key", "value")

	secretCmd.AddCommand(secretListCmd, secretCreateCmd, secretDeleteCmd, secretGetCmd, secretUpdateCmd)
	rootCmd.AddCommand(secretCmd)
}
