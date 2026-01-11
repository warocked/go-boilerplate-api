package handlers

import (
	"context"
	"go-boilerplate-api/internal/api/db"
	"go-boilerplate-api/shared/helpers"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Health(ctx *fiber.Ctx) error {
	healthCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	checks := fiber.Map{}
	httpStatus := fiber.StatusOK

	if db.DB != nil {
		err := db.HealthCheckGORM(healthCtx)
		poolStats, statsErr := db.GetPoolStats()
		
		if err != nil {
			checks["postgres"] = fiber.Map{
				"status": "error",
				"message": err.Error(),
			}
			httpStatus = fiber.StatusServiceUnavailable
		} else {
			postgresCheck := fiber.Map{"status": "ok"}
			if statsErr == nil && poolStats != nil {
				postgresCheck["stats"] = fiber.Map{
					"max_open_connections": poolStats.MaxOpenConnections,
					"open_connections":     poolStats.OpenConnections,
					"in_use":               poolStats.InUse,
					"idle":                 poolStats.Idle,
					"wait_count":           poolStats.WaitCount,
					"wait_duration":        poolStats.WaitDuration,
					"max_idle_closed":      poolStats.MaxIdleClosed,
					"max_idle_time_closed": poolStats.MaxIdleTimeClosed,
					"max_lifetime_closed":  poolStats.MaxLifetimeClosed,
				}
			}
			checks["postgres"] = postgresCheck
		}
	}

	if db.RedisClient != nil {
		err := db.HealthCheckRedis(healthCtx)
		if err != nil {
			checks["redis"] = fiber.Map{
				"status": "error",
				"message": err.Error(),
			}
			if httpStatus == fiber.StatusOK {
				httpStatus = fiber.StatusServiceUnavailable
			}
		} else {
			checks["redis"] = fiber.Map{"status": "ok"}
		}
	}

	return helpers.SendSuccess(ctx, httpStatus, checks, "Health check completed")
}
