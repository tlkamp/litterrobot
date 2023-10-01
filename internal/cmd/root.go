package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	lrc "github.com/tlkamp/litter-api/v2/pkg/client"
)

func newLitterRobotCmd(c *lrc.Client, version string) *cobra.Command {
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
			viper.ReadInConfig() //nolint:errcheck
			return nil
		},
	}

	cmd.AddCommand(
		newVersionCmd(version),
		newConfigCmd(),
	)

	return cmd
}

func Execute(v string) error {
	var client *lrc.Client
	return newLitterRobotCmd(client, v).Execute()
}
