package usecases

import (
	"context"

	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/Archisman-Mridha/chat-service/internal/types"
)

type (
	GetChatMessagesArgs struct {
		ChatID            int32
		PaginationOptions types.PaginationOptions
	}

	GetChatMessagesOutput struct {
		Messages []types.Message
	}
)

func (u *Usecases) GetChatMessages(ctx context.Context, args *GetChatMessagesArgs) (*GetChatMessagesOutput, error) {
	panic(constants.UNIMPLEMENTED)
}
