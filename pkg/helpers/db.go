package helpers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"

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

func GetRedisConnParams() (string, int64, string, string) {
	username := ""
	password := ""
	host := "localhost"
	var port int64 = 6379

	if hostEnv := os.Getenv("REDIS_HOST"); hostEnv != "" {
		host = hostEnv
	}

	if usernameEnv := os.Getenv("REDIS_USERNAME"); usernameEnv != "" {
		username = usernameEnv
	}

	if passwordEnv := os.Getenv("REDIS_PASSWORD"); passwordEnv != "" {
		password = passwordEnv
	}

	if portEnv := os.Getenv("REDIS_PORT"); portEnv != "" {
		port, _ = strconv.ParseInt(portEnv, 10, 64)
	}

	return host, port, username, password
}

func GetRedisConnection(db int) *redis.Client {
	redisHost, redisPort, redisUsername, redisPassword := GetRedisConnParams()
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword,
		Username: redisUsername,
		DB:       db,
	})
}

func CreateDatabase(connectionStr, dbName string) error {
	conn, err := CreateDatabaseConnection(connectionStr)
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), "CREATE SCHEMA "+dbName)

	if err != nil {
		return err
	}

	return conn.Close(context.Background())
}
