package utils

import (
	"context"
	"log/slog"
	"os"

	"github.com/Archisman-Mridha/go-microservice-template/pkg/assert"
)

// Returns the value of the given environment variable.
//
// Panics if the environment variable isn't set.
func GetEnv(name string) string {
	envValue, envFound := os.LookupEnv(name)
	assert.Assert(context.Background(), envFound, "Env not found", slog.String("env", name))

	return envValue
}
