package client

import (
	"context"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Register or login with following commands.",
	Long:  `Register or login with following commands.`,
}

var loginAuthCmd = &cobra.Command{
	Use:   "login -u <USERNAME> -p <PASSWORD>",
	Short: "Login with following commands.",
	Long:  `Login with following commands.`,
	RunE:  AuthLogin,
}

var registerAuthCmd = &cobra.Command{
	Use:   "register -u <USERNAME> -p <PASSWORD>",
	Short: "Register with following commands.",
	Long:  `Register with following commands.`,
	RunE:  AuthRegister,
}

// AuthLogin authenticats user on server.
func AuthLogin(cmd *cobra.Command, args []string) error {
	log.Debug().Msgf("running %s", cmd.Name())
	u, p, err := getUsernameAndPassword(cmd)
	if err != nil {
		return nil
	}

	ctx := cmd.Root().Context()
	c := ctx.Value(CtxKey("c")).(*GRPCAgent)

	req := &pb.LoginRequest{
		Login:    u,
		Password: p,
	}

	resp, err := c.authClient.service.Login(context.Background(), req)
	if err != nil {
		return err
	}
	c.SaveToken(resp.Token)

	log.Info().Msgf("Login is successful. Server response: %+v", resp)
	return nil
}

// AuthRegister registers user on server.
func AuthRegister(cmd *cobra.Command, args []string) error {
	log.Debug().Msgf("running %s", cmd.Name())

	u, p, err := getUsernameAndPassword(cmd)
	if err != nil {
		return nil
	}

	ctx := cmd.Root().Context()
	c := ctx.Value(CtxKey("c")).(*GRPCAgent)

	req := &pb.RegisterRequest{
		Login:    u,
		Password: p,
	}
	resp, err := c.authClient.service.Register(context.Background(), req)
	if err != nil {
		log.Err(err).Msg("Error")
		return err
	}

	c.SaveToken(resp.Token)

	log.Info().Msgf("User is registered successfully. Server response: %+v", resp)
	return nil
}

// getUsernameAndPassword oarses username and password from flags.
func getUsernameAndPassword(cmd *cobra.Command) (string, string, error) {
	username, err := cmd.Flags().GetString("username")
	if err != nil {
		return "", "", err
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}

func init() {
	log.Info().Msgf("Calling auth")
	var u, pw string
	loginAuthCmd.Flags().StringVarP(&u, "username", "u", "", "Username (required if password is set)")
	loginAuthCmd.Flags().StringVarP(&pw, "password", "p", "", "Password (required if username is set)")
	loginAuthCmd.MarkFlagsRequiredTogether("username", "password")

	registerAuthCmd.Flags().StringVarP(&u, "username", "u", "", "Username (required if password is set)")
	registerAuthCmd.Flags().StringVarP(&pw, "password", "p", "", "Password (required if username is set)")
	registerAuthCmd.MarkFlagsRequiredTogether("username", "password")

	authCmd.AddCommand(loginAuthCmd, registerAuthCmd)

	rootCmd.AddCommand(authCmd)
}
