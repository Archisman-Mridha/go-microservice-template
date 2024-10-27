package types

import (
	"context"
	"time"

	"github.com/Archisman-Mridha/chat-service/pkg/healthcheck"
	"github.com/google/uuid"
)

type (
	Database interface {
		healthcheck.Healthcheckable

		GetChats(ctx context.Context, userID int32, paginationOptions PaginationOptions) ([]*Chat, error)
		GetMessages(ctx context.Context, chatID int32, paginationOptions PaginationOptions) ([]*Message, error)
		CreateChat(ctx context.Context) (chatID int32, err error)
		CreateMessage(ctx context.Context, id uuid.UUID, chatID int32, message string, senderID int32, sentAt time.Time) error
	}

	KVStore interface {
		healthcheck.Healthcheckable

		// Set key to hold the string value. If key already holds a value, it is overwritten, regardless
		// of its type. Any previous time to live associated with the key is discarded on successful SET
		// operation.
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error

		// Get the value of key. An error is returned if the value stored at key is not a string,
		// because GET only handles string values.
		Get(ctx context.Context, key string) (*string, error)

		// Removes the specified keys. A key is ignored if it does not exist.
		Del(ctx context.Context, keys ...string) error
	}
)
