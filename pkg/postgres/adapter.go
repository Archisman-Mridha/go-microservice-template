package postgres

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/Archisman-Mridha/chat-service/internal/types"
	"github.com/Archisman-Mridha/chat-service/pkg/logger"
	"github.com/Archisman-Mridha/chat-service/pkg/postgres/sql/generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/lib/pq"
)

type DatabaseAdapter struct {
	connection *pgx.Conn
	queries    *generated.Queries
}

func NewDatabaseAdapter(ctx context.Context, postgresURL string) types.Database {
	connection, err := pgx.Connect(ctx, postgresURL)
	if err != nil {
		log.Fatalf("Failed creating Postgres connection : %v", err)
	}

	if err := connection.Ping(ctx); err != nil {
		log.Fatalf("Failed pinging Postgres : %v", err)
	}
	slog.InfoContext(ctx, "Connected with Postgres")

	queries := generated.New(connection)

	return &DatabaseAdapter{
		connection,
		queries,
	}
}

func (d *DatabaseAdapter) Healthcheck() error {
	ctx := context.Background()
	if err := d.connection.Ping(ctx); err != nil {
		return fmt.Errorf("failed pinging Postgres : %v", err)
	}
	return nil
}

func (d *DatabaseAdapter) GetChats(ctx context.Context, userID int32, paginationOptions types.PaginationOptions) ([]*types.Chat, error) {
	getChatsParams := generated.GetChatsParams{
		UserID: userID,
		Offset: paginationOptions.Offset,
		Limit:  paginationOptions.Limit,
	}
	rows, err := d.queries.GetChats(ctx, getChatsParams)
	if err != nil {
		slog.ErrorContext(ctx, "Failed executing GetChats query", slog.Any("params", getChatsParams), logger.Error(err))
		return nil, err
	}

	chats := []*types.Chat{}
	for _, row := range rows {
		lastMessageID, err := uuid.FromBytes(row.LastMessageID.Bytes[:])
		if err != nil {
			slog.ErrorContext(ctx, "Failed converting from pgtype.UUID to uuid.UUID", slog.Any("pgtype_uuid", row.LastMessageID), logger.Error(err))
			return nil, err
		}

		lastMessage := &types.Message{
			ID:       lastMessageID,
			ChatID:   row.ID,
			SenderID: row.LastMessageSenderID,
			Message:  row.LastMessage,
			SentAt:   row.LastMessageSentAt.Time,
		}

		chats = append(chats, &types.Chat{
			ID:          row.ID,
			WithUserID:  row.WithUserID,
			LastMessage: lastMessage,
		})
	}
	return chats, nil
}

func (d *DatabaseAdapter) GetMessages(ctx context.Context, chatID int32, paginationOptions types.PaginationOptions) ([]*types.Message, error) {
	getMessagesParams := generated.GetMessagesParams{
		ChatID: chatID,
		Offset: paginationOptions.Offset,
		Limit:  paginationOptions.Limit,
	}
	rows, err := d.queries.GetMessages(ctx, getMessagesParams)
	if err != nil {
		slog.ErrorContext(ctx, "Failed executing GetMessages query", slog.Any("params", getMessagesParams), logger.Error(err))
		return nil, err
	}

	messages := []*types.Message{}
	for _, row := range rows {
		id, err := uuid.FromBytes(row.ID.Bytes[:])
		if err != nil {
			slog.ErrorContext(ctx, "Failed converting from pgtype.UUID to uuid.UUID", slog.Any("pgtype_uuid", row.ID), logger.Error(err))
			return nil, err
		}

		messages = append(messages, &types.Message{
			ID:       id,
			ChatID:   chatID,
			SenderID: row.SenderID,
			Message:  row.Message,
			SentAt:   row.SentAt.Time,
		})
	}
	return messages, nil
}

func (d *DatabaseAdapter) CreateChat(ctx context.Context) (int32, error) {
	chatID, err := d.queries.CreateChat(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed executing CreateChat query", logger.Error(err))
		return 0, err
	}
	return chatID, nil
}

func (d *DatabaseAdapter) CreateMessage(ctx context.Context, id uuid.UUID, chatID int32, message string, senderID int32, sentAt time.Time) error {
	createMessageParams := generated.CreateMessageParams{
		ID:       pgtype.UUID{Bytes: id, Valid: true},
		ChatID:   chatID,
		Message:  message,
		SenderID: senderID,
		SentAt: pgtype.Timestamptz{
			Time: sentAt,
		},
	}
	if err := d.queries.CreateMessage(ctx, createMessageParams); err != nil {
		createMessageParams.Message = "" // For privacy reasons, we'll not log the actual message.
		slog.ErrorContext(ctx, "Failed executing CreateMessage query", slog.Any("params", createMessageParams), logger.Error(err))

		return err
	}
	return nil
}
