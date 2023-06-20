package cmd

import (
	"github.com/erfanmomeniii/user-management/internal/app"
	"github.com/spf13/cobra"
)

var (
	configPath string
	a          *app.App
)

var rootCmd = &cobra.Command{
	Use:               "user-management",
	Short:             "A service that can handle and save user information.",
	Long:              `user-management is a service that can handle and save user information.`,
	PersistentPreRunE: preRun,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&configPath, "config", "c", "", "Config file path",
	)

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(migrateCmd)
}

func preRun(_ *cobra.Command, _ []string) (err error) {
	a, err = app.New(configPath)
	if err != nil {
		panic(err)
	}

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}
