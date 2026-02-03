package helpers

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func GetEnvAsInt(key string, defaultValue uint16) uint16 {
    valueStr := GetEnv(key, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return uint16(value)
    }
    return defaultValue
}