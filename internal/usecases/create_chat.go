package usecases

import (
	"context"

	"github.com/Archisman-Mridha/chat-service/constants"
)

type (
	CreateChatArgs struct {
		UserIDs []int32
	}

	CreateChatOutput struct {
		ChatID int32
	}
)

func (u *Usecases) CreateChat(ctx context.Context, args *CreateChatArgs) (*CreateChatOutput, error) {
	panic(constants.UNIMPLEMENTED)
}
