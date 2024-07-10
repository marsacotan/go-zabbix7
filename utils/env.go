package utils

import (
	"log"
	"os"
)

func getEnv(envVar string) string {
	value, exists := os.LookupEnv(envVar)
	if !exists {
		log.Fatalf("Environment variable %s not set", envVar)
	}
	return value
}

func LoadEnvVarToken(envVar string) string {
	return getEnv(envVar)
}
