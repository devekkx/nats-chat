package logger

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/devekkx/nats-chat/packages/config"
	"github.com/devekkx/nats-chat/packages/logger"
	"golang.org/x/sys/unix"

	natsclient "github.com/devekkx/nats-chat/packages/nats"
)

func main() {
	cfg := config.Load()

	log := logger.New(
		cfg.ServiceName,
		cfg.Environment,
	)
	client, err := natsclient.New(cfg.NATSUrl)

	if err != nil {
		log.Error("NATS connection failed", slog.String("error", err.Error()))
		return
	}

	defer client.Close()

	client.Subscribe("chat.>", func(data []byte) {
		log.Info("event received", slog.String("payload", string(data)))
	})

	log.Info("Logger service started")

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, unix.SIGTERM)

	<-stop
	log.Info("Shutting down gracefully...")

}
