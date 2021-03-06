package helpers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"os"
	"strconv"

	"github.com/SSH-Management/utils/v2"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/db/config"
)

const (
	connectionOptions = "application_name=SSHManagementTest&sslmode=disable&search_path=%s"

	connectionStringFmt = "postgresql://%s:%s@%s:%d/%s?%s"
)

func findMigrationsDir(workingDir string) (string, error) {
	for entries, err := os.ReadDir(workingDir); err == nil; {
		for _, entry := range entries {
			if entry.IsDir() && entry.Name() == "migrations" {
				return workingDir + "/" + entry.Name(), nil
			}
		}

		workingDir, err = utils.GetAbsolutePath(workingDir + "/..")

		if err != nil {
			return "", err
		}

		entries, err = os.ReadDir(workingDir)
	}

	return "", errors.New("migrations dir not found")
}

func SetupDatabase() (*gorm.DB, func()) {
	var bytes [16]byte
	_, _ = rand.Read(bytes[:])

	dbRandomIndex := hex.EncodeToString(bytes[:])

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	if database == "" {
		database = "ssh_management"
	}

	if username == "" {
		username = "postgres"
	}

	if password == "" {
		password = "postgres"
	}

	if host == "" {
		host = "localhost"
	}

	var port int64 = 5432

	if portStr != "" {
		port, _ = strconv.ParseInt(portStr, 10, 32)
	}

	schema := fmt.Sprintf("ssh_management_%s", dbRandomIndex)

	connectionString := fmt.Sprintf(connectionStringFmt, username, password, host, port, database, fmt.Sprintf(connectionOptions, schema))

	if err := CreateDatabase(connectionString, schema); err != nil {
		panic(err)
	}

	clean := func() {
		conn, err := CreateDatabaseConnection(connectionString)
		if err != nil {
			panic(err)
		}

		if _, err = conn.Exec(context.Background(), "DROP SCHEMA "+schema+" CASCADE"); err != nil {
			panic(err)
		}

		if err = conn.Close(context.Background()); err != nil {
			panic(err)
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		clean()
		panic(err)
	}

	migrationsDir, err := findMigrationsDir(wd)
	if err != nil {
		clean()
		panic(err)
	}

	migrations, err := migrate.New("file://"+migrationsDir, connectionString)
	if err != nil {
		clean()
		panic(err)
	}

	if err = migrations.Up(); err != nil && err != migrate.ErrNoChange {
		clean()
		panic(err)
	}

	gormDb, err := db.GetDbConnection(config.Config{
		Username:        username,
		Password:        password,
		Database:        database,
		Schema:          schema,
		Host:            host,
		Port:            int(port),
		MaxIdleTime:     0,
		ConnMaxLifetime: 0,
		ConnMaxIdle:     1,
		ConnMaxOpen:     10,
	})

	if err != nil {
		clean()
		panic(err)
	}

	return gormDb, clean
}
