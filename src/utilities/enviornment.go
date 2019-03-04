package utilities

import (
	"os"
)

func getEnvironmentVariable(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
