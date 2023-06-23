package cmd

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"

	Migrate "github.com/golang-migrate/migrate"
	MigrateMySQL "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/cobra"

	"github.com/erfanmomeniii/user-management/internal/app"
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
	a, err := app.New(configPath)
	if !(strings.HasPrefix(path, "/")) {
		wd, err := os.Getwd()
		if err != nil {
			a.Logger.Fatal("cannot get working directory", zap.Error(err))
		}

		path, err = filepath.Abs(filepath.Join(wd, path))
		if err != nil {
			a.Logger.Fatal("cannot get absolute path", zap.Error(err))
		}
	}

	a.Logger.Info("migrating...", zap.String("path", path))

	driver, err := MigrateMySQL.WithInstance(a.Database.DB, &MigrateMySQL.Config{MigrationsTable: table})
	if err != nil {
		a.Logger.Fatal("cannot setup mysql driver", zap.Error(err))
	}

	m, err := Migrate.NewWithDatabaseInstance("file://"+path, a.Config.Database.Name, driver)

	if err != nil {
		a.Logger.Fatal("cannot instantiate migrate", zap.Error(err))
	}

	if steps == 0 {
		err = m.Up()
	} else {
		err = m.Steps(steps)
	}

	if err != nil {
		if err == Migrate.ErrNoChange {
			a.Logger.Info("no change applied to database")
		} else {
			a.Logger.Fatal("cannot migrate", zap.Error(err))
		}
	} else {
		a.Logger.Info("migration done successfully")
	}
}
