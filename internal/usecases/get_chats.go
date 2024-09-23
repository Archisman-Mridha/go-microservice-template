package usecases

import (
	"context"

	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/Archisman-Mridha/chat-service/internal/types"
)

type (
	GetChatsArgs struct {
		UserID int32
	}

	GetChatsOutput struct {
		Chat []types.Chat
	}
)

func (u *Usecases) GetChats(ctx context.Context, args *GetChatsArgs) (*GetChatsOutput, error) {
	panic(constants.UNIMPLEMENTED)
}
