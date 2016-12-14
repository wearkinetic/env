package env

import (
	"log"
	"os"
)

// Parse returns the list of environment values we care about
func Parse(keys ...string) (*map[string]string, error) {

	parsed := map[string]string{}

	for _, key := range keys {

		value := os.Getenv(key)
		if value == "" {
			log.Println("Please provide with the following environment variables")
			for _, key_ := range keys {
				log.Println(key_)
			}
			return &parsed, errNoEnv(key)
		}
		parsed[key] = value
	}

	return &parsed, nil

}
