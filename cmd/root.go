package cmd

import (
	"github.com/spf13/cobra"
	"user-management/internal/app"
)

var configPath string

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

func preRun(_ *cobra.Command, _ []string) error {
	if err := app.Init(configPath); err != nil {
		panic(err)
	}

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}
