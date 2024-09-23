package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	"github.com/Archisman-Mridha/chat-service/internal/types"
	_ "github.com/lib/pq"
)

type DatabaseAdapter struct {
	connection *sql.DB
}

func NewDatabaseAdapter(postgresURL string) types.Database {
	connection, err := sql.Open("postgres", postgresURL)
	if err != nil {
		log.Fatalf("failed creating Postgres connection : %v", err)
	}

	if err := connection.Ping(); err != nil {
		log.Fatalf("failed pinging Postgres : %v", err)
	}
	slog.Info("Connected with Postgres")

	return &DatabaseAdapter{
		connection,
	}
}

func (a *DatabaseAdapter) Healthcheck() error {
	if err := a.connection.Ping(); err != nil {
		return fmt.Errorf("failed pinging Postgres : %v", err)
	}
	return nil
}
