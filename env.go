package env

import "os"

// Parse returns the list of environment values we care about
func Parse(keys ...string) (*map[string]string, error) {

	parsed := map[string]string{}

	for _, key := range keys {

		value := os.Getenv(key)
		if value == "" {
			return &parsed, errNoEnv(key)
		}
		parsed[key] = value
	}

	return &parsed, nil

}
