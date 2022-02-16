package helpers

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func CreateDatabaseConnection(connectionStr string) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(connectionStr)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateDatabase(connectionStr, dbName string) error {
	conn, err := CreateDatabaseConnection(connectionStr)
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), "CREATE DATABASE "+dbName)

	if err != nil {
		return err
	}

	return conn.Close(context.Background())
}
