package helpers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase(connectionStr, dbName string) error {
	sqlDB, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return err
	}

	_, err = sqlDB.Exec("CREATE DATABASE " + dbName + " CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")

	if err != nil {
		return err
	}

	if err = sqlDB.Close(); err != nil {
		return err
	}

	return nil
}
