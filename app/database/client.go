package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

// DB is a db.
type DB struct{}

// NewDB creates a DB.
func NewDB() *DB {
	return &DB{}
}

// Conn connects to database.
func (d *DB) Conn() (*redis.Client, error) {
	dataSourceName := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	conn := redis.NewClient(&redis.Options{
		Addr:     dataSourceName,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := conn.Ping().Result()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
