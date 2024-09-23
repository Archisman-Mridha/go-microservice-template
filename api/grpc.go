package api

import (
	"context"

	"github.com/Archisman-Mridha/chat-service/api/proto/generated"
	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/Archisman-Mridha/chat-service/internal/usecases"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatServiceGRPCAPI struct {
	generated.UnimplementedChatServiceServer
	usecases *usecases.Usecases
}

func NewChatServiceGRPCAPI(usecases *usecases.Usecases) *ChatServiceGRPCAPI {
	return &ChatServiceGRPCAPI{
		usecases: usecases,
	}
}

func (c *ChatServiceGRPCAPI) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (c *ChatServiceGRPCAPI) GetChats(ctx context.Context, request *generated.GetChatsRequest) (*generated.GetChatsResponse, error) {

	panic(constants.UNIMPLEMENTED)

}

func (c *ChatServiceGRPCAPI) CreateChat(ctx context.Context, request *generated.CreateChatRequest) (*generated.CreateChatResponse, error) {

	panic(constants.UNIMPLEMENTED)

}

func (c *ChatServiceGRPCAPI) GetChatMessages(ctx context.Context, request *generated.GetChatMessagesRequest) (*generated.GetChatMessagesResponse, error) {

	panic(constants.UNIMPLEMENTED)

}

func (c *ChatServiceGRPCAPI) Chat(generated.ChatService_ChatServer) error {

	panic(constants.UNIMPLEMENTED)

}
