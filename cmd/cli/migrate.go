package cli

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/cmd/command"
	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/db/config"
	"github.com/SSH-Management/server/pkg/helpers"
)

func migrateCommand(migrations embed.FS) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "migrate",
		Short:             "Migrate database to latest version",
		PersistentPreRunE: command.LoadConfig,
		RunE:              migrateDatabase(migrations),
	}

	flags := cmd.PersistentFlags()

	flags.BoolP("create-database", "c", false, "Create Database in PostgreSQL Server")
	flags.BoolP("create-default-roles", "d", false, "Insert Default roles")

	return cmd
}

func createDatabase(cfg config.Config) error {
	c := cfg.Clone()
	c.Database = "(empty)"
	return helpers.CreateDatabase(c.FormatConnectionString(), cfg.Database)
}

func migrateDatabase(migrations embed.FS) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		flags := cmd.PersistentFlags()

		shouldCreateDatabase, err := flags.GetBool("create-database")
		if err != nil {
			return err
		}

		shouldInsertDefaultRoles, err := flags.GetBool("create-default-roles")
		if err != nil {
			return err
		}

		_ = shouldInsertDefaultRoles

		migrationsFS, err := iofs.New(migrations, "migrations")
		if err != nil {
			return err
		}

		var cfg config.Config

		c := command.GetContainer()
		defer func(c *container.Container) {
			_ = c.Close()
		}(c)

		if err := c.Config.Sub("database").Unmarshal(&cfg); err != nil {
			return err
		}

		if shouldCreateDatabase {
			if err := createDatabase(cfg); err != nil {
				return err
			}
		}

		connStr := fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=%s",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database,
			cfg.SSLMode,
			cfg.TimeZone,
		)

		m, err := migrate.NewWithSourceInstance("iofs", migrationsFS, connStr)
		if err != nil {
			return err
		}

		if err = m.Up(); err != nil {
			return err
		}

		m.Close()

		return nil
	}
}
