package middlewares

import (
	"go-boilerplate-api/internal/api/config"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"

	"github.com/rs/zerolog"

	"github.com/gofiber/fiber/v2"
)

func SetupLogger(app *fiber.App) {
	switch config.LOG_LEVEL {
	case config.LOG_LEVEL_DEBUG:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Stack().Logger()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
		Fields: []string{fiberzerolog.FieldIPs, fiberzerolog.FieldLatency, fiberzerolog.FieldStatus,
			fiberzerolog.FieldMethod, fiberzerolog.FieldURL, fiberzerolog.FieldError},
	}))
}
