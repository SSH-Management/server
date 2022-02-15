package helpers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	"github.com/SSH-Management/utils/v2"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/db/config"
)

const (
	connectionOptions = "charset=utf8mb4&checkConnLiveness=true&collation=utf8mb4_general_ci&interpolateParams=true&loc=UTC&multiStatements=true&parseTime=true"

	connectionStringFmt = "mysql://%s:%s@tcp(%s:%d)/ssh_management_%s?%s"
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

	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	portStr := os.Getenv("MYSQL_PORT")

	if username == "" {
		username = "root"
	}

	if password == "" {
		password = "password"
	}

	if host == "" {
		host = "localhost"
	}

	var port int64 = 3306

	if portStr != "" {
		port, _ = strconv.ParseInt(portStr, 10, 32)
	}

	dbName := fmt.Sprintf("ssh_management_%s", dbRandomIndex)

	connectionString := fmt.Sprintf(connectionStringFmt, username, password, host, port, dbRandomIndex, connectionOptions)

	sqlDBConnectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/?%s", username, password, host, port, connectionOptions)

	if err := CreateMySQLDatabase(sqlDBConnectionStr, dbName); err != nil {
		panic(err)
	}

	clean := func() {
		sqlDB, err := sql.Open("mysql", sqlDBConnectionStr)
		if err != nil {
			panic(err)
		}

		if _, err = sqlDB.Exec("DROP DATABASE " + dbName); err != nil {
			panic(err)
		}

		if err = sqlDB.Close(); err != nil {
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

	if err := migrations.Up(); err != nil {
		clean()
		panic(err)
	}

	gormDb, err := db.GetDbConnection(config.Config{
		Driver:          "mysql",
		Username:        username,
		Password:        password,
		Database:        dbName,
		Collation:       "utf8mb4_general_ci",
		Host:            fmt.Sprintf("%s:%d", host, port),
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
