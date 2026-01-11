package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func LoadAllConfig() error {
	// Load env file
	err := LoadEnvFile()
	if err != nil {
		return err
	}

	LOG_LEVEL, err = determineLogLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return err
	}

	PORT = os.Getenv("PORT")
	IS_PROD = os.Getenv("IS_PROD") == "true"

	SECRET_KEY = os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		SECRET_KEY = "qweasd123"
	}

	expireRaw := os.Getenv("TOKEN_EXPIRE_TIME")
	if expireRaw != "" {
		TOKEN_TTL, err = time.ParseDuration(expireRaw)
		if err != nil {
			return fmt.Errorf("error parsing token expire time duration, %w", err)
		}
	} else {
		TOKEN_TTL = time.Hour * 5
	}

	ALLOWED_ORIGINS = os.Getenv("ALLOWED_ORIGINS")

	REDIS_KEYS_TTL, err = time.ParseDuration(os.Getenv("REDIS_KEYS_TTL"))
	if err != nil {
		REDIS_KEYS_TTL = time.Hour * 24 * 7
	}

	DATABASE_URL = os.Getenv("DATABASE_URL")
	REDIS_URL = os.Getenv("REDIS_URL")

	timezone := os.Getenv("TIMEZONE")
	if timezone != "" {
		TIMEZONE = timezone
	}

	bodyLimitStr := os.Getenv("REQUEST_BODY_LIMIT_MB")
	if bodyLimitStr != "" {
		var bodyLimitMB int
		_, err := fmt.Sscanf(bodyLimitStr, "%d", &bodyLimitMB)
		if err == nil && bodyLimitMB > 0 {
			REQUEST_BODY_LIMIT = bodyLimitMB * 1024 * 1024
		}
	}

	return nil
}

func LoadEnvFile() error {
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		return godotenv.Load(".env")
	}

	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		envPath := filepath.Join(exeDir, ".env")
		if _, err := os.Stat(envPath); !os.IsNotExist(err) {
			return godotenv.Load(envPath)
		}
	}

	return nil
}

func determineLogLevel(logLevel string) (LOG_LEVEL_TYPE, error) {
	switch logLevel {
	case "":
		return LOG_LEVEL_DEBUG, nil
	case "info":
		return LOG_LEVEL_INFO, nil
	case "warn":
		return LOG_LEVEL_WARN, nil
	case "debug":
		return LOG_LEVEL_DEBUG, nil
	case "error":
		return LOG_LEVEL_ERROR, nil
	case "fatal":
		return LOG_LEVEL_FATAL, nil
	default:
		return LOG_LEVEL_NOTFOUND, fmt.Errorf("invalid log level")
	}
}
