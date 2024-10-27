package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	Chat struct {
		ID,
		WithUserID int32
		LastMessage *Message
	}

	Message struct {
		ID uuid.UUID
		ChatID,
		SenderID,
		ReceiverID int32
		Message string
		SentAt  time.Time
	}

	IncomingMessage struct {
		WorkerID string // ID of the worker which sent this incoming message for the user.
		Message  Message
	}
)
