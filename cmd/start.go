package cmd

import (
	"github.com/spf13/cobra"
	httpServer "user-management/internal/http/server"
	"user-management/internal/shutdown"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: startFunc,
}

func startFunc(_ *cobra.Command, _ []string) {
	httpServer.Serve()
	shutdown.Wait()
}
