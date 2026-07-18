package config

import "os"

type Config struct {
	ServiceName string
	Environment string
	NATSUrl     string
}

func Load() Config {

	return Config{
		ServiceName: getEnv("SERVICE_NAME", "unknwon-service"),
		Environment: getEnv("ENVIRONMENT", "development"),
		NATSUrl:     getEnv("NATS_URL", "nats://localhost:4222"),
	}
}

func getEnv(
	key string,
	fallback string,
) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
