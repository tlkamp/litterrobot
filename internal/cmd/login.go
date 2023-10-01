package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	lrc "github.com/tlkamp/litter-api/v2/pkg/client"
)

func newLoginCmd(c *lrc.Client) *cobra.Command {
	return &cobra.Command{
		Use:          "login",
		Short:        "log into the litter robot 3",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			c = lrc.New(viper.GetString("email"), viper.GetString("password"))

			if err := c.Login(context.Background()); err != nil {
				return err
			}

			viper.Set("token", c.Token())
			return viper.WriteConfig()
		},
	}
}
