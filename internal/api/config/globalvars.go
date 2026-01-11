package config

import "time"

type LOG_LEVEL_TYPE int8

const (
	LOG_LEVEL_NOTFOUND LOG_LEVEL_TYPE = iota - 1
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_DEBUG
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL
)

var (
	PORT            = "8080"
	IS_PROD         = false
	LOG_LEVEL       = LOG_LEVEL_DEBUG
	SECRET_KEY      = "#!@$%^&*()_019baea7-823a-7da1-8f9c-7d4677aa76d4"
	ALLOWED_ORIGINS = ""
	REDIS_KEYS_TTL  = time.Hour * 24 * 7
	TOKEN_TTL       = time.Hour * 5
	S3BUCKETNAME    = "testbucket"
	TIMEZONE        = "Asia/Manila"

	DATABASE_URL = ""
	REDIS_URL    = ""

	REQUEST_BODY_LIMIT = 50 * 1024 * 1024
)
