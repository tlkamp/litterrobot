package cmd

import "github.com/spf13/cobra"

func newVersionCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "print the version information",
		Run: func(c *cobra.Command, _ []string) {
			c.Println(version)
		},
	}
}
