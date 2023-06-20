package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A start the application",
	Long:  "A start the app server",
	Run:   startFunc,
}

func startFunc(_ *cobra.Command, _ []string) {
	a.Start()

	ctx := a.Wait()

	a.Shutdown(ctx)
}
