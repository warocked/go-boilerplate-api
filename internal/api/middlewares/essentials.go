package middlewares

import (
	"log"
	"os"
	"strings"
	"time"

	"go-boilerplate-api/internal/api/config"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"
)

func SetupMiddlewaresEssentials(app *fiber.App) {
	SetupMiddlewareRecover(app)
	SetupMiddlewareRequestID(app)
	SetupMiddlewareHelmet(app)
	SetupMiddlewareCORS(app)
	SetupMiddlewareRateLimiter(app)
	SetupMiddlewareCompress(app)
	SetupMiddlewareFiberZerolog(app)
}

// SetupMiddlewareRecover recovers from panics and prevents server crashes
func SetupMiddlewareRecover(app *fiber.App) {
	app.Use(recover.New(recover.Config{
		EnableStackTrace: !config.IS_PROD, // Only show stack traces in development
	}))
}

// SetupMiddlewareRequestID adds unique request ID to each request
func SetupMiddlewareRequestID(app *fiber.App) {
	app.Use(requestid.New())
}

// SetupMiddlewareCompress compresses response body
func SetupMiddlewareCompress(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // Use best speed for high throughput
	}))
}

// SetupMiddlewareHelmet sets security headers
func SetupMiddlewareHelmet(app *fiber.App) {
	app.Use(helmet.New(helmet.Config{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		ReferrerPolicy:        "no-referrer",
		CrossOriginEmbedderPolicy: "require-corp",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginResourcePolicy: "cross-origin",
		HSTSMaxAge:            31536000, // 1 year
		HSTSExcludeSubdomains: false,
		ContentSecurityPolicy: "default-src 'self'",
	}))
}

// SetupMiddlewareCORS configures CORS based on environment
func SetupMiddlewareCORS(app *fiber.App) {
	origins := config.ALLOWED_ORIGINS
	originList := []string{}

	if origins == "" {
		if config.IS_PROD {
			// In production, require explicit configuration
			log.Fatalf("FATAL: ALLOWED_ORIGINS must be set in production. Cannot use default origins when IS_PROD=true")
		} else {
			// Development default
			origins = "http://localhost:3000"
		}
	}

	// Parse comma-separated origins
	for _, origin := range strings.Split(origins, ",") {
		trimmed := strings.TrimSpace(origin)
		if trimmed != "" && trimmed != "*" {
			// When AllowCredentials is true, cannot use wildcard "*"
			originList = append(originList, trimmed)
		}
	}

	// Validate: cannot use wildcard with credentials
	if origins == "*" {
		if config.IS_PROD {
			log.Fatalf("FATAL: Cannot use wildcard '*' for ALLOWED_ORIGINS in production when AllowCredentials is enabled")
		} else {
			log.Printf("WARNING: Wildcard '*' not allowed with AllowCredentials, defaulting to http://localhost:3000")
			originList = []string{"http://localhost:3000"}
		}
	}

	// If no valid origins after parsing, use development default
	if len(originList) == 0 {
		if config.IS_PROD {
			log.Fatalf("FATAL: No valid origins configured. ALLOWED_ORIGINS must be set in production")
		} else {
			originList = []string{"http://localhost:3000"}
		}
	}

	// When AllowCredentials is true, cannot use wildcard "*"
	// Must specify exact origins
	corsConfig := cors.Config{
		AllowOrigins:     strings.Join(originList, ","),
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS,HEAD",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,Accept-Language,Content-Length",
		AllowCredentials: true,
		MaxAge:           3600, // 1 hour
	}

	app.Use(cors.New(corsConfig))
}

// SetupMiddlewareRateLimiter prevents abuse and DDoS attacks
func SetupMiddlewareRateLimiter(app *fiber.App) {
	// Different limits for production vs development
	maxRequests := 100 // requests per window
	if config.IS_PROD {
		maxRequests = 100 // Adjust based on your needs
	}

	app.Use(limiter.New(limiter.Config{
		Max:        maxRequests,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit by IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "rate_limit_exceeded",
				"message": "Too many requests, please try again later",
			})
		},
		SkipSuccessfulRequests: false,
		SkipFailedRequests:     false,
	}))
}

// SetupMiddlewareFiberZerolog sets up structured logging
func SetupMiddlewareFiberZerolog(app *fiber.App) {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))
}
