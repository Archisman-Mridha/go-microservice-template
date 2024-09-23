package utils

import (
	"flag"
	"log"
	"os"
	"strings"
)

type GetFlagOrEnvValueFn = func(f *flag.Flag)

// Usage : flagSet.VisitAll(getFlagOrEnvValue("CHAT_MICROSERVICE_"))
func CreateGetFlagOrEnvValueFn(envPrefix string) GetFlagOrEnvValueFn {

	// If a flag isn't set, then we try to get its value from the corresponding environment variable.
	// Panics, if the flag and environment variable aren't set and there's no default flag value.
	//
	// NOTE : If the flag name is config-file-path, then the corresponding environment variable is
	// CHAT_MICROSERVICE_CONFIG_FILE_PATH. CHAT_MICROSERVICE_ here is the env-prefix.
	getFlagOrEnvValueFn := func(f *flag.Flag) {
		if len(f.Value.String()) > 0 {
			return
		}

		// Since the flag is not set, we'll try to get the value from the corresponding environment
		// variable.
		envName := envPrefix + strings.ReplaceAll(strings.ToUpper(f.Name), "-", "_")
		envValue, envFound := os.LookupEnv(envName)
		if envFound {
			if err := f.Value.Set(envValue); err != nil {
				log.Fatalf("Failed setting value of flag %v to value of environment variable %s", f.Name, envName)
			}
			return
		}

		if len(f.DefValue) == 0 {
			log.Fatalf("Neither flag %s nor environment variable %s was set", f.Name, envName)
		}
	}

	return getFlagOrEnvValueFn
}
