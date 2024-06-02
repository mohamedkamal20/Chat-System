package redis

import (
	"context"
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
)

var redisClient *goredislib.Client

func initRedisClient(ctx context.Context, db int) {
	if redisClient != nil {
		return // Client already initialized
	}
	redisClient = goredislib.NewClient(&goredislib.Options{
		Addr:     "redis:6379", // Redis container hostname and port
		Password: "",           // No password set
		DB:       db,           // Use specified DB
	})

	// Ping the Redis server to ensure connectivity
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error initializing Redis client:", err)
	} else {
		fmt.Println("Redis client initialized successfully:", pong)
	}
}

func SetMessages() {
	ctx := context.Background()
	db := 0 // Set the desired Redis DB number
	initRedisClient(ctx, db)

	err := redisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
	}
}

func GetMessages(email string) /* ([]map[string]interface{}, error) */ {
	ctx := context.Background()
	db := 0 // Set the desired Redis DB number
	initRedisClient(ctx, db)

	//messages, err := redisClient.LRange(ctx, email, 0, -1).Result()
	//if err != nil {
	//	fmt.Println("Error getting key:", err)
	//	return nil, err
	//}
	//return messages, nil
}
