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
