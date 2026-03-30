package app

import (
	grpcapp "grpc/internal/app/grpc"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger,
	grpcPort int,
	storagePath string) *App {

	//TODO: инициализировать хранилище (storage
	//TODO: init service
	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
