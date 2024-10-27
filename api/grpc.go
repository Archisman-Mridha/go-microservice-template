package api

import (
	"context"
	"io"
	"log/slog"

	"github.com/Archisman-Mridha/chat-service/api/proto/generated"
	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/Archisman-Mridha/chat-service/internal/types"
	"github.com/Archisman-Mridha/chat-service/internal/usecases"
	"github.com/Archisman-Mridha/chat-service/pkg/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatServiceGRPCAPI struct {
	generated.UnimplementedChatServiceServer
	usecases *usecases.Usecases
	kvStore  types.KVStore
}

func NewChatServiceGRPCAPI(usecases *usecases.Usecases, kvStore types.KVStore) *ChatServiceGRPCAPI {
	return &ChatServiceGRPCAPI{
		usecases: usecases,
		kvStore:  kvStore,
	}
}

func (c *ChatServiceGRPCAPI) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (c *ChatServiceGRPCAPI) GetChats(ctx context.Context, request *generated.GetChatsRequest) (*generated.GetChatsResponse, error) {
	output, err := c.usecases.GetChats(ctx, &usecases.GetChatsArgs{
		UserID: request.UserId,
		PaginationOptions: types.PaginationOptions{
			Offset: request.PaginationOptions.Offset,
			Limit:  request.PaginationOptions.Limit,
		},
	})
	if err != nil {
		return nil, status.Error(getErrorStatusCode(err), err.Error())
	}

	chats := []*generated.Chat{}
	for _, chat := range output.Chats {
		chats = append(chats, &generated.Chat{
			Id:         chat.ID,
			WithUserId: chat.WithUserID,

			LastMessage: &generated.Message{
				Id:       chat.LastMessage.ID.String(),
				ChatId:   chat.LastMessage.ChatID,
				Message:  chat.LastMessage.Message,
				SenderId: chat.LastMessage.SenderID,
				SentAt:   timestamppb.New(chat.LastMessage.SentAt),
			},
		})
	}
	return &generated.GetChatsResponse{Chats: chats}, nil
}

func (c *ChatServiceGRPCAPI) GetMessages(ctx context.Context, request *generated.GetMessagesRequest) (*generated.GetMessagesResponse, error) {
	output, err := c.usecases.GetMessages(ctx, &usecases.GetMessagesArgs{
		ChatID: request.ChatId,
		PaginationOptions: types.PaginationOptions{
			Offset: request.PaginationOptions.Offset,
			Limit:  request.PaginationOptions.Limit,
		},
	})
	if err != nil {
		return nil, status.Error(getErrorStatusCode(err), err.Error())
	}

	messages := []*generated.Message{}
	for _, message := range output.Messages {
		messages = append(messages, &generated.Message{
			Id:       message.ID.String(),
			ChatId:   message.ChatID,
			Message:  message.Message,
			SenderId: message.SenderID,
			SentAt:   timestamppb.New(message.SentAt),
		})
	}
	return &generated.GetMessagesResponse{Messages: messages}, nil
}

func (c *ChatServiceGRPCAPI) Chat(stream generated.ChatService_ChatServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()

	handshakeDone := false

	// When a message sent by the user gets successfully processed, it gets pushed to this Go
	// channel. It then gets picked up and an acknowledgement gets constructed of off it. The
	// acknowledgement is then pushed back to the client, so that it doesn't retry sending that
	// message.
	acknowledgementsToUserChan := make(chan *generated.Message)
	defer close(acknowledgementsToUserChan)

	// If any error occurs in the go routines, then that error is pushed to this Go channel. The error
	// is then picked up and an appropriate gRPC error is sent to the client, thus, closing the stream.
	errChan := make(chan error)
	defer close(errChan)

	// Handling messages sent by the user.
	go func() {
		for {
			oneofRequest, err := stream.Recv()
			if err == io.EOF { // Client has closed the stream.
				cancel()
			}
			if err != nil {
				slog.ErrorContext(ctx, "Unexpected error occurred", logger.Error(err))
				continue
			}

			switch request := oneofRequest.Request.(type) {
			case *generated.ChatStreamRequest_Handshake:
				if handshakeDone {
					slog.WarnContext(ctx, "Received handshake request more than once")
					continue
				}

				if err := c.chatStreamHandshakeRequestHandler(ctx, request.Handshake, stream); err != nil {
					errChan <- err
					return
				}
				handshakeDone = true

			case *generated.ChatStreamRequest_Message:
				// Handshake must have been done by this point. The client needs to send a handshake request
				// first, before sending message requests.
				if !handshakeDone {
					errChan <- constants.ErrHandshakeNotDone
					return
				}

				c.chatMessageRequestHandler(ctx, request.Message)
				acknowledgementsToUserChan <- request.Message
			}
		}
	}()

	for {
		select {
		// Pushing an acknowledgement back to the user, that a message he / she sent before has been
		// successfully processed by the server.
		case message := <-acknowledgementsToUserChan:
			err := stream.Send(&generated.ChatStreamResponse{
				Response: &generated.ChatStreamResponse_Acknowledgement{
					Acknowledgement: &generated.Acknowledgement{
						MessageId: message.Id,
						ChatId:    message.ChatId,
					},
				},
			})
			if err != nil {
				slog.ErrorContext(ctx, "Failed sending acknowledgement back to the user")

				// Push the message back to the acknowledgementsToUserChan again, in order to retry sending
				// the acknowledgement.
				acknowledgementsToUserChan <- message
			}

		case err := <-errChan:
			return status.Error(getErrorStatusCode(err), err.Error())

		case <-ctx.Done():
			return nil
		}
	}
}

func (c *ChatServiceGRPCAPI) chatStreamHandshakeRequestHandler(
	ctx context.Context,
	handshake *generated.Handshake,
	stream generated.ChatService_ChatServer,
) error {
	// Save the userID -> pod IP mapping in the KV store.
	// We close the stream if the operation fails.
	if err := c.kvStore.Set(ctx, string(handshake.UserId), constants.InstanceID, 0); err != nil {
		return constants.ErrHandshakeFailed
	}

	constants.UserIDToStream[handshake.UserId] = stream

	return nil
}

func (c *ChatServiceGRPCAPI) chatMessageRequestHandler(ctx context.Context, message *generated.Message) {
	messageID, err := uuid.FromBytes([]byte(message.Id))
	if err != nil {
		slog.ErrorContext(ctx, "Failed parsing message ID as UUID", slog.String("id", message.Id), logger.Error(err))
		return
	}

	err = c.usecases.CreateMessage(ctx, &usecases.CreateMessageArgs{
		ID:       messageID,
		ChatID:   message.ChatId,
		Message:  message.Message,
		SenderID: message.SenderId,
		SentAt:   message.SentAt.AsTime(),
	})
	if err != nil {
		slog.ErrorContext(ctx, "Unexpected error occurred", logger.Error(err))
		return
	}
}

func (c *ChatServiceGRPCAPI) ReceiveIncomingMessage(ctx context.Context, incomingMessage *generated.IncomingMessage) (*emptypb.Empty, error) {
	stream, ok := constants.UserIDToStream[incomingMessage.ReceiverId]
	if !ok {
		slog.ErrorContext(ctx, "Received incoming message for user who isn't connected to this server", slog.Int("user_id", int(incomingMessage.ReceiverId)))

		err := constants.ErrUserNotConnectedToServer
		return nil, status.Error(getErrorStatusCode(err), err.Error())
	}

	err := stream.Send(&generated.ChatStreamResponse{
		Response: &generated.ChatStreamResponse_Message{
			Message: incomingMessage.Message,
		},
	})
	if err != nil {
		slog.ErrorContext(ctx, "Failed sending incoming message to the user", logger.Error(err))

		err = constants.ErrServer
		return nil, status.Error(getErrorStatusCode(err), err.Error())
	}

	return &emptypb.Empty{}, nil
}

// Returns the gRPC status code based on the given error.
func getErrorStatusCode(err error) (code codes.Code) {
	switch err {
	case constants.ErrUserNotConnectedToServer:
	case constants.ErrHandshakeNotDone:
		code = codes.FailedPrecondition

	case constants.ErrHandshakeFailed:
	case constants.ErrServer:
	default:
		code = codes.Internal
	}
	return
}
