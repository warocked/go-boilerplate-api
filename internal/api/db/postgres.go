package db

import (
	"context"
)

func InitPostgres(ctx context.Context, databaseURL string) error {
	return InitGORM(ctx, databaseURL)
}

func ClosePostgres() error {
	return CloseGORM()
}

func HealthCheck(ctx context.Context) error {
	return HealthCheckGORM(ctx)
}