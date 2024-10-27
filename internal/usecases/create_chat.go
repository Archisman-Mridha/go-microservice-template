package usecases

import (
	"context"

	"github.com/Archisman-Mridha/chat-service/constants"
)

type CreateChatOutput struct {
	ChatID int32
}

func (u *Usecases) CreateChat(ctx context.Context) (*CreateChatOutput, error) {
	chatID, err := u.db.CreateChat(ctx)
	if err != nil {
		return nil, constants.ErrServer
	}

	return &CreateChatOutput{ChatID: chatID}, nil
}
