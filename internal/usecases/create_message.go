package usecases

import (
	"context"
	"time"

	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/google/uuid"
)

type (
	CreateMessageArgs struct {
		ID       uuid.UUID
		ChatID   int32
		Message  string
		SenderID int32
		SentAt   time.Time
	}

	CreateMessageOutput struct {
		MessageID int32
	}
)

func (u *Usecases) CreateMessage(ctx context.Context, args *CreateMessageArgs) error {
	err := u.db.CreateMessage(ctx, args.ID, args.ChatID, args.Message, args.SenderID, args.SentAt)
	if err != nil {
		return constants.ErrServer
	}

	return nil
}
