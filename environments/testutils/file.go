package testutils

import (
	"encoding/json"
	"fmt"
	"os"
)

func SetEnvsFromFile(filepath string) (cleanup func()) {
	if filepath == "" {
		return func() {}
	}

	envs, _, err := loadJsonFromFile[map[string]any](filepath)
	if err != nil {
		panic(err)
	}

	return envSetter(envs)
}

func loadJsonFromFile[T any](filePath string) (T, []byte, error) {
	var t T
	data, err := os.ReadFile(filePath)
	if err != nil {
		return t, data, err
	}

	err = json.Unmarshal(data, &t)
	return t, data, err
}

func envSetter(envs map[string]any) (cleanup func()) {
	originalEnvs := map[string]any{}

	for name, value := range envs {
		if originalValue, ok := os.LookupEnv(name); ok {
			originalEnvs[name] = originalValue
		}
		_ = os.Setenv(name, fmt.Sprint(value))
	}

	return func() {
		for name := range envs {
			origValue, has := originalEnvs[name]
			if has {
				_ = os.Setenv(name, fmt.Sprint(origValue))
			} else {
				_ = os.Unsetenv(name)
			}
		}
	}
}
