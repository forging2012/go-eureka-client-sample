package cmd

import (
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server as eureka client",
	Long:  `Server as eureka client.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func run() error {
	return nil
}
