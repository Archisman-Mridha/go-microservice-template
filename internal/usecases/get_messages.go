package usecases

import (
	"context"

	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/Archisman-Mridha/chat-service/internal/types"
)

type (
	GetMessagesArgs struct {
		ChatID            int32
		PaginationOptions types.PaginationOptions
	}

	GetMessagesOutput struct {
		Messages []*types.Message
	}
)

func (u *Usecases) GetMessages(ctx context.Context, args *GetMessagesArgs) (*GetMessagesOutput, error) {
	messages, err := u.db.GetMessages(ctx, args.ChatID, args.PaginationOptions)
	if err != nil {
		return nil, constants.ErrServer
	}

	return &GetMessagesOutput{Messages: messages}, nil
}
