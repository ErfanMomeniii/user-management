package cmd

import (
	"github.com/spf13/cobra"
)

var (
	steps           int
	migrationsPath  string
	migrationsTable string
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	Long:  "Run migration on the database",
	Run:   migrate,
}

func init() {
	migrateCmd.Flags().StringVarP(&migrationsPath, "migrations-path", "m", "./migrations",
		"path to migrations directory")
	migrateCmd.Flags().StringVarP(&migrationsTable, "migrations-table", "t", "schema_migrations",
		"database table holding migrations")
	migrateCmd.Flags().IntVarP(&steps, "steps", "n", 0,
		"number of steps to migrate. positive steps for up and negative steps for down. zero to upgrade all.")
}

func migrate(_ *cobra.Command, _ []string) {
}
