package config

import (
	"os"
	"strconv"
	"strings"
)

// struct that holds everything we need to talk to the minio
type MinIOConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
}

type Config struct {
	MinIO MinIOConfig
}

func getEnvString(key string, defaultVal string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return defaultVal
	}
	return val
}

func getEnvBool(key string, defaultVal bool) bool {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return defaultVal
	}

	parsed, err := strconv.ParseBool(raw)
	if err != nil {
		return defaultVal
	}
	return parsed
}

// returns a config object the rest of the app uses
func Load() Config {
	var cfg Config

	cfg.MinIO = MinIOConfig{
		Endpoint:  getEnvString("MINIO_ENDPOINT", "localhost:9000"),
		AccessKey: getEnvString("MINIO_ACCESS_KEY", "minioadmin"),
		SecretKey: getEnvString("MINIO_SECRET_KEY", "minioadmin"),
		Bucket:    getEnvString("MINIO_BUCKET", "media"),
		UseSSL:    getEnvBool("MINIO_USE_SSL", false),
	}

	return cfg
}
