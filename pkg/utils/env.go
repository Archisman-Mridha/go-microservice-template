package utils

import (
	"log"
	"os"
)

// Returns the value of the given environment variable.
// Panics if the environment variable isn't set.
func GetEnv(name string) string {
	envValue, envFound := os.LookupEnv(name)
	if !envFound {
		log.Fatalf("Environment variable %s not found", name)
	}
	return envValue
}
