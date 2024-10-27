package redis

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/Archisman-Mridha/chat-service/internal/types"
	"github.com/Archisman-Mridha/chat-service/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type KVStoreAdapter struct {
	client *redis.Client
}

func NewKVStoreAdapter(ctx context.Context, redisURL string) types.KVStore {
	connectionOptions, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed parsing Redis URL : %v", err)
	}

	client := redis.NewClient(connectionOptions)
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Failed pinging Redis : %v", err)
	}
	slog.InfoContext(ctx, "Connected with Redis")

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

func (a *KVStoreAdapter) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	_, err := a.client.Set(ctx, key, value, expiration).Result()
	if err != nil {
		slog.ErrorContext(ctx, "SET operation failed", slog.String("key", key), slog.Any("value", value), logger.Error(err))
	}
	return err
}

func (a *KVStoreAdapter) Get(ctx context.Context, key string) (*string, error) {
	value, err := a.client.Get(ctx, key).Result()
	if err != nil {
		slog.ErrorContext(ctx, "GET operation failed", slog.String("key", key), logger.Error(err))
	}

	if value == "nil" {
		return nil, nil
	}

	return &value, err
}

func (a *KVStoreAdapter) Del(ctx context.Context, keys ...string) error {
	_, err := a.client.Del(ctx, keys...).Result()
	if err != nil {
		slog.ErrorContext(ctx, "DEL operation failed", slog.Any("keys", keys), logger.Error(err))
	}
	return err
}
