package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	// RedisClient is the global Redis client with connection pooling
	RedisClient *redis.Client
)

// InitRedis initializes the Redis client with optimal settings for high performance
func InitRedis(ctx context.Context, redisURL string) error {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	// Optimize pool settings for high throughput and millions of requests
	opt.PoolSize = 100              // Maximum number of socket connections
	opt.MinIdleConns = 10           // Minimum number of idle connections
	opt.MaxRetries = 3              // Maximum number of retries for failed commands
	opt.DialTimeout = time.Second * 5
	opt.ReadTimeout = time.Second * 3
	opt.WriteTimeout = time.Second * 3
	opt.PoolTimeout = time.Second * 4 // Timeout for getting connection from pool
	opt.ConnMaxIdleTime = time.Minute * 5  // Close connections after remaining idle

	// Create Redis client
	RedisClient = redis.NewClient(opt)

	// Verify the connection by pinging Redis
	err = RedisClient.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to ping Redis: %w", err)
	}

	return nil
}

// CloseRedis closes the Redis client gracefully
func CloseRedis() error {
	if RedisClient != nil {
		return RedisClient.Close()
	}
	return nil
}

// GetRedis returns the global Redis client
func GetRedis() *redis.Client {
	return RedisClient
}

// HealthCheck checks if the Redis connection is healthy
func HealthCheckRedis(ctx context.Context) error {
	if RedisClient == nil {
		return fmt.Errorf("Redis client is not initialized")
	}
	return RedisClient.Ping(ctx).Err()
}