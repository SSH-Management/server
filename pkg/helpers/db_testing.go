package helpers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	"github.com/SSH-Management/utils/v2"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/db/config"
)

const (
	connectionOptions = "application_name=SSHManagementTest&sslmode=disable"

	connectionStringFmt = "postgresql://%s:%s@%s:%d/ssh_management_%s?%s"
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

	dbName := fmt.Sprintf("ssh_management_%s", dbRandomIndex)

	connectionString := fmt.Sprintf(connectionStringFmt, username, password, host, port, dbRandomIndex, connectionOptions)

	sqlDBConnectionStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/?%s", username, password, host, port, connectionOptions)

	if err := CreateDatabase(sqlDBConnectionStr, dbName); err != nil {
		panic(err)
	}

	clean := func() {
		conn, err := CreateDatabaseConnection(sqlDBConnectionStr)
		if err != nil {
			panic(err)
		}

		if _, err = conn.Exec(context.Background(), "DROP DATABASE "+dbName+" WITH (FORCE)"); err != nil {
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

	if err = migrations.Up(); err != nil {
		clean()
		panic(err)
	}

	gormDb, err := db.GetDbConnection(config.Config{
		Username:        username,
		Password:        password,
		Database:        dbName,
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
