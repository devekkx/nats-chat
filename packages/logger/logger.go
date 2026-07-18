package logger

import (
	"log/slog"
	"os"
)

func New(service, environment string) *slog.Logger {
	var handler slog.Handler

	if environment == "production" {
		handler = slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		)
	} else {
		handler = slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug},
		)
	}

	logger := slog.New(handler)

	return logger.With(slog.String("service", service))
}
