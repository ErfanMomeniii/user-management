package cmd

import (
	Migrate "github.com/golang-migrate/migrate/v4"
	MigrateMySQL "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
	"user-management/internal/config"
	"user-management/internal/database"
	"user-management/internal/log"
)

var (
	steps int
	path  string
	table string
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	Long:  "Run migration on the database",
	Run:   migrate,
}

func init() {
	migrateCmd.Flags().StringVarP(&path, "path", "p", "./migrations",
		"path to migrations directory")
	migrateCmd.Flags().StringVarP(&table, "table", "t", "schema_migrations",
		"database table holding migrations")
	migrateCmd.Flags().IntVarP(&steps, "steps", "s", 0,
		"number of steps to migrate. positive steps for up and negative steps for down. zero to upgrade all.")
}

func migrate(_ *cobra.Command, _ []string) {
	if path == "" {
		log.L.Fatal("the path (p) argument is required")
	}

	if !(strings.HasPrefix(path, "/")) {
		wd, err := os.Getwd()
		if err != nil {
			log.L.Fatal("cannot get working directory", zap.Error(err))
		}

		path, err = filepath.Abs(filepath.Join(wd, path))
		if err != nil {
			log.L.Fatal("cannot get absolute path", zap.Error(err))
		}
	}

	log.L.Info("migrating...", zap.String("path", path))

	driver, err := MigrateMySQL.WithInstance(database.DB.DB, &MigrateMySQL.Config{MigrationsTable: table})
	if err != nil {
		log.L.Fatal("cannot setup mysql driver", zap.Error(err))
	}

	m, err := Migrate.NewWithDatabaseInstance("file://"+path, config.C.Database.Name, driver)
	if err != nil {
		log.L.Fatal("cannot instantiate migrate", zap.Error(err))
	}

	if steps == 0 {
		err = m.Up()
	} else {
		err = m.Steps(steps)
	}

	if err != nil {
		if err == Migrate.ErrNoChange {
			log.L.Info("no change applied to database")
		} else {
			log.L.Fatal("cannot migrate", zap.Error(err))
		}
	} else {
		log.L.Info("migration done successfully")
	}
}
