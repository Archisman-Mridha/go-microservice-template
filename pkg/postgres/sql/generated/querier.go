// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package generated

import (
	"context"
)

type Querier interface {
	CreateChat(ctx context.Context) error
	CreateChatMessage(ctx context.Context, arg CreateChatMessageParams) error
	CreateChatParticipant(ctx context.Context, arg CreateChatParticipantParams) error
	GetChatMessages(ctx context.Context, arg GetChatMessagesParams) ([]GetChatMessagesRow, error)
	GetChats(ctx context.Context, arg GetChatsParams) ([]GetChatsRow, error)
}

var _ Querier = (*Queries)(nil)