package cmd

import (
	"github.com/spf13/cobra"
	httpServer "user-management/internal/http/server"
	"user-management/internal/shutdown"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A start the application",
	Long:  "A start the app server",
	Run:   startFunc,
}

func startFunc(_ *cobra.Command, _ []string) {
	httpServer.Serve()
	shutdown.Wait()
}
