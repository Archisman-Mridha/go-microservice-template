package logger

import (
	"log/slog"
	"os"
)

func InitLogger() {
	logger := slog.New(
		withContextualSlogAttributesHandler(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})))
	slog.SetDefault(logger)
}

func Error(err error) slog.Attr {
	return slog.Any("err", err)
}
