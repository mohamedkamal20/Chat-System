package redis

import (
	"Chat-System/models"
	"context"
	"encoding/json"
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"log"
	"time"
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

	cacheKey := fmt.Sprintf("messages:%s", email)
	// Serialize the list of messages to JSON
	jsonData, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("Error serializing messages:", err)
		return
	}
	err = redisClient.Set(ctx, cacheKey, jsonData, time.Minute*10).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
	}
}

func GetMessages(email string) ([]map[string]interface{}, error) {
	ctx := context.Background()
	db := 0
	initRedisClient(ctx, db)

	cacheKey := fmt.Sprintf("messages:%s", email)
	val, err := redisClient.Get(ctx, cacheKey).Result()
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

func InvalidateCacheMessages(email string) {
	ctx := context.Background()
	db := 0
	initRedisClient(ctx, db)
	cacheKey := fmt.Sprintf("messages:%s", email)
	redisClient.Del(ctx, cacheKey)
}

func InvalidateCacheUser(email string) {
	ctx := context.Background()
	db := 0
	initRedisClient(ctx, db)
	cacheKey := fmt.Sprintf("user_data:%s", email)
	redisClient.Del(ctx, cacheKey)
}

func SetUser(email string, user models.User) {
	ctx := context.Background()
	db := 0
	initRedisClient(ctx, db)

	cacheKey := fmt.Sprintf("user_data:%s", email)

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshaling user data:", err)
		return
	}
	err = redisClient.Set(ctx, cacheKey, userJson, time.Hour*24).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
	}
}

func GetUser(email string) (*models.User, error) {
	ctx := context.Background()
	db := 0
	initRedisClient(ctx, db)

	cacheKey := fmt.Sprintf("user_data:%s", email)
	val, err := redisClient.Get(ctx, cacheKey).Result()
	if err == goredislib.Nil {
		return nil, nil // Cache miss
	} else if err != nil {
		fmt.Println("Error getting key:", err)
		return nil, err
	}

	// Deserialize the JSON string back to a list of messages

	var user models.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		fmt.Println("Error deserializing messages:", err)
		return nil, err
	}

	return &user, nil
}
