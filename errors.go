package env

import (
	"errors"
	"fmt"
)

// errNoEnv is returned when an environment value is missing
func errNoEnv(key string) error {
	return errors.New(fmt.Sprintf("%s - Key: %s", ERR_NO_ENV, key))
}
