package cmd

import (
	"github.com/carllhw/go-eureka-client-sample/pkg/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server as eureka client",
	Long:  `Server as eureka client.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func run() {
	server.Start()
}
