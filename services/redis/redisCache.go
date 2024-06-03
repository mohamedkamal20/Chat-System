package redis

import (
	"context"
	"encoding/json"
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

func SetMessages(email string, messages []map[string]interface{}) {
	ctx := context.Background()
	db := 0 // Set the desired Redis DB number
	initRedisClient(ctx, db)

	// Serialize the list of messages to JSON
	jsonData, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("Error serializing messages:", err)
		return
	}
	err = redisClient.Set(ctx, email, jsonData, 0).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
	}
}

func GetMessages(email string) ([]map[string]interface{}, error) {
	ctx := context.Background()
	db := 0 // Set the desired Redis DB number
	initRedisClient(ctx, db)

	val, err := redisClient.Get(ctx, email).Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
		return nil, err
	}

	// Deserialize the JSON string back to a list of messages
	var retrievedMessages []map[string]interface{}
	err = json.Unmarshal([]byte(val), &retrievedMessages)
	if err != nil {
		fmt.Println("Error deserializing messages:", err)
		return nil, err
	}
	return retrievedMessages, nil
}
