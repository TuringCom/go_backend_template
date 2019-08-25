package util

import "os"

// GetEnv retrieves environment variable or uses fallback value
func GetEnv(envKey, fallback string) string {
	if value, ok := os.LookupEnv(envKey); ok {
		return value
	}
	return fallback
}
