package types

import "time"

type (
	Chat struct {
		ID,
		WithUserID int32
		LastMessage Message
	}

	Message struct {
		ID,
		ChatID,
		SenderID int32
		Message string
		SentAt  time.Time
	}
)
