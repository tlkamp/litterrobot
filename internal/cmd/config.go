package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const showSecrets = "show-secrets"

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "config",
		Short:        "Manage local configuration for the CLI",
		SilenceUsage: true,
	}

	cmd.AddCommand(
		newConfigInitCmd(),
		newConfigSetCmd(),
		newConfigGetCmd(),
		newConfigShowCmd(),
	)

	return cmd
}

func newConfigInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "init",
		Short:        "initialize config file",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			viper.Set("email", "email@example.com")
			viper.Set("password", "password")
			return viper.SafeWriteConfig()
		},
	}
}

func newConfigSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "set config values",
		Args:  cobra.ExactArgs(2),
		RunE: func(c *cobra.Command, args []string) error {
			key := args[0]
			value := args[1]

			viper.Set(key, value)
			return viper.WriteConfig()
		},
	}
}

func newConfigGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "get",
		Short:        "get config values",
		SilenceUsage: true,
		Args:         cobra.MaximumNArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			show, err := c.Flags().GetBool(showSecrets)
			if err != nil {
				return err
			}

			if show {
				c.Println(viper.Get(args[0]))
			} else {
				c.Println(getRedacted(args[0], viper.GetString(args[0])))
			}

			return nil
		},
	}

	cmd.Flags().BoolP(showSecrets, "s", false, "print secret values")

	return cmd
}

func newConfigShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "show",
		Short:        "print configuration",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			show, err := c.Flags().GetBool(showSecrets)
			if err != nil {
				return err
			}

			if err := viper.ReadInConfig(); err != nil {
				return err
			}

			for k, v := range viper.AllSettings() {
				if show {
					c.Println(k+":", v)
					continue
				}

				c.Println(k+":", getRedacted(k, viper.GetString(k)))
			}
			return nil
		},
	}

	cmd.Flags().BoolP(showSecrets, "s", false, "Print secret values")

	return cmd
}

func getRedacted(k, v string) string {
	switch k {
	case "password", "token":
		return "*******"
	default:
		return v
	}
}
