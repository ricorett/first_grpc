package main

import (
	"fmt"
	"grpc/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO: инициализировать объект конфига
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//TODO: инициализировать логгер
	log := setupLogger(cfg.Env)
	log.Info("starting server", slog.String("env", cfg.Env))
	//TODO: инициализировать приложение
	//TODO: инициализировать gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
