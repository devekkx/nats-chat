package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	Service string
	Env     string
}

func New(cfg Config) *slog.Logger {
	var handler slog.Handler

	if cfg.Env == "production" {
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

	return logger.With(slog.String("service", cfg.Service))
}
