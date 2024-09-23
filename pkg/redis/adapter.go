package redis

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/Archisman-Mridha/chat-service/internal/types"
	"github.com/redis/go-redis/v9"
)

type KVStoreAdapter struct {
	client *redis.Client
}

func NewKVStoreAdapter(redisURL string) types.KVStore {
	connectionOptions, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("failed parsing Redis URL : %v", err)
	}

	client := redis.NewClient(connectionOptions)
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("failed pinging Redis : %v", err)
	}
	slog.Info("Connected with Redis")

	return &KVStoreAdapter{
		client,
	}
}

func (a *KVStoreAdapter) Healthcheck() error {
	if _, err := a.client.Ping(context.Background()).Result(); err != nil {
		return fmt.Errorf("failed pinging Redis : %v", err)
	}
	return nil
}
