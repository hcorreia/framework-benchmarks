package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func getEnvOrFallback(key string, fallback string) string {
	value := os.Getenv(key)

	if strings.TrimSpace(value) == "" {
		return fallback
	}

	return value
}

func getEnvOrPanic(key string) string {
	value := os.Getenv(key)

	if strings.TrimSpace(value) == "" {
		log.Fatalf("Missing %s environment variable", key)
	}

	return value
}

var Env = struct {
	Addr          string
	ChaosEndpoint string
}{
	Addr:          getEnvOrFallback("ADDR", ":8000"),
	ChaosEndpoint: getEnvOrPanic("CHAOS_ENDPOINT"),
}
