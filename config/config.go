package config

type Config struct {
	GRPCServerPort    int `yaml:"gRPCServerPort" default:"4000"`
	MetricsServerPort int `yaml:"metricsServerPort" default:"5000"`

	RedisURL    string `yaml:"redisURL"`
	PostgresURL string `yaml:"postgresURL"`
	JaegerURL   string `yaml:"jaegerURL" validate:"required"`
}
