package utils

import (
	"os"

	"github.com/argonsecurity/go-environments/models"
	funk "github.com/thoas/go-funk"
)

func SetValueIfEnvExist(configuration *models.Configuration, envName string, path string) {
	if envValue, ok := os.LookupEnv(envName); ok {
		funk.Set(configuration, envValue, path)
	}
}
