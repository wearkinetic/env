package env

import "os"

// Parse returns the list of environment values we care about
func Parse(keys ...string) (*map[string]string, error) {

	parsed := map[string]string{}
	missingKeys := []string{}

	for _, key := range keys {

		value := os.Getenv(key)
		if value == "" {
			missingKeys = append(missingKeys, key)
		}
		parsed[key] = value
	}

	if len(missingKeys) > 0 {
		return &parsed, errNoEnv(missingKeys...)
	}

	return &parsed, nil

}
