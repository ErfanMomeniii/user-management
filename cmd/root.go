package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "user-management",
	Short: "A service that can handle and save user information.",
	Long:  `user-management is a service that can handle and save user information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&configPath, "config", "c", "", "Config file path",
	)

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(migrateCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
