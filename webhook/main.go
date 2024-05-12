package main

import (
	"context"
	"log"
	"os"

	redisClient "webhook/redis"

	"github.com/go-redis/redis/v8" // Make sure to use the correct version
)

func main() {
	// Create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initializing the Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})

	// Create a channel to act as the queue
	webhookQueue := make(chan redisClient.WebhookPayload, 100) // Setting buffer size as 100

	// Subscribe to the "transactions" channel
	err := redisClient.Subscribe(ctx, client, webhookQueue)

	if err != nil {
		log.Println("Error:", err)
	}

	select {}
}
