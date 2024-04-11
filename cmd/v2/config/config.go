package config

import (
	"os"
	"time"
)

const (
	GCThresholdPercent = 0.7 // GCThresholdPercent is the threshold for garbage collection

	GCLimit = 1024 * 1024 * 1024 // GCLimit is the limit for garbage collection

	MB = 1024 * 1024 // MB is the number of bytes in a megabyte

	SentryFlushTimeout = 2 * time.Second // SentryFlushTimeout is the timeout for flushing sentry

	NonceLength = 16 // NonceLength is the length of the nonce : 16 bytes * 8 bits/byte = 128 bits
)

func GetConfigPath(isDevelopment bool) string {
	if isDevelopment {
		return "./cmd/v2/config/config-local"
	}

	return "./cmd/v2/config/config-prod"
}

func IsProduction() bool {
	return os.Getenv("APP_ENV") != "dev"
}

func IsDevelopment() bool {
	return os.Getenv("APP_ENV") == "dev"
}

func GetCallbackURL() string {
	return os.Getenv("CALLBACK_URL")
}
