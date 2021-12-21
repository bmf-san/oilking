package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

// Redis is a db.
type Redis struct{}

// NewRedis creates a Redis.
func NewRedis() *Redis {
	return &Redis{}
}

// Conn connects to redis.
func (d *Redis) Conn() (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := conn.Ping().Result()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
