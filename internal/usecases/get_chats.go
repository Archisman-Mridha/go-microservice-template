package usecases

import (
	"context"

	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/Archisman-Mridha/chat-service/internal/types"
)

type (
	GetChatsArgs struct {
		UserID            int32
		PaginationOptions types.PaginationOptions
	}

	GetChatsOutput struct {
		Chats []*types.Chat
	}
)

func (u *Usecases) GetChats(ctx context.Context, args *GetChatsArgs) (*GetChatsOutput, error) {
	chats, err := u.db.GetChats(ctx, args.UserID, args.PaginationOptions)
	if err != nil {
		return nil, constants.ErrServer
	}

	return &GetChatsOutput{Chats: chats}, nil
}
