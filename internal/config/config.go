package config

import (
	"github.com/Archisman-Mridha/go-microservice-template/pkg/connectors"
	"github.com/Archisman-Mridha/go-microservice-template/pkg/utils"
)

type (
	Config struct {
		DevMode      bool `yaml:"devMode" default:"False"`
		DebugLogging bool `yaml:"debugLogging" default:"False"`

		ServerPort int `yaml:"serverPort" default:"4000" validate:"gt=0"`

		JWTSigningKey string `yaml:"jwtSigningKey" validate:"required,notblank"`

		Postgres                 connectors.NewPostgresConnectorArgs `yaml:"postgres" validate:"required"`
		Redis                    connectors.NewRedisConnectorArgs    `yaml:"redis" validate:"required"`
		OpenTelemetryCollectrURL string                              `yaml:"openTelemetryCollectorURL" validate:"required,notblank"`
		Flagsmith                utils.GetOpenFeatureClientArgs      `yaml:"flagsmith" validate:"required"`
	}
)
