package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"gopkg.in/yaml.v3"
)

// Parses the config file at the given path.
// NOTE : An environment variable (if found), will override the corresponding config.
func ParseConfigFile[T any](configFilePath string) (*T, error) {
	configFileContents, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed reading config file %s : %v", configFilePath, err)
	}
	config := new(T)
	if err := yaml.Unmarshal(configFileContents, config); err != nil {
		return nil, fmt.Errorf("failed YAML unmarshalling config file : %v", err)
	}

	// Confirm that required fields are set.
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(config); err != nil {
		log.Fatalf("Config validation failed : %v", err)
	}

	// Populate optional fields with corresponding default values.
	defaults.SetDefaults(config)

	return config, nil
}
