package utils

import (
	"log"
	"os"
)

func GetEnv(envVar string) string {
	value, exists := os.LookupEnv(envVar)
	if !exists {
		log.Fatalf("Environment variable %s not set", envVar)
	}
	return value
}
