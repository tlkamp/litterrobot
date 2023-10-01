package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	lrc "github.com/tlkamp/litter-api/v2/pkg/client"
)

func newLitterRobotCmd(c *lrc.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "litterrobot",
		Short: "Control your Litter Robot 3 Connect",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			dir, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			configPath := path.Join(dir, ".litterrobot")

			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			viper.AddConfigPath(configPath)
			viper.ReadInConfig()
			return nil
		},
	}

	return cmd
}

func Execute() error {
	var client *lrc.Client
	return newLitterRobotCmd(client).Execute()
}