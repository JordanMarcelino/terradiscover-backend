package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jordanmarcelino/terradiscover-backend/internal/config"
	"github.com/jordanmarcelino/terradiscover-backend/internal/logger"
	"github.com/jordanmarcelino/terradiscover-backend/internal/server"
)

func main() {
	cfg := config.InitConfig()
	logger.SetZerologLogger(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	srv := server.NewHttpServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
