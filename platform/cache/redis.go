package cache

import (
	"log"
	"os"
	"strconv"

	"backend/pkg/utils"

	// "github.com/go-redis/redis/v8"
	"github.com/go-redis/redis"
)

// RedisConnection func for connect to Redis server.
// func RedisConnection() (*redis.Client, error) {
func RedisConnection() *redis.Client {
	// Define Redis database number.
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		// return nil, err
		log.Panic("Failed connect to redis")
	}

	// Set Redis options.
	options := &redis.Options{
		Addr:     redisConnURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}

	// return redis.NewClient(options), nil
	return redis.NewClient(options)
}
