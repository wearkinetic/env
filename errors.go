package env

import (
	"errors"
	"fmt"
	"strings"
)

// errNoEnv is returned when an environment value is missing
func errNoEnv(keys ...string) error {
	k := strings.Join(keys, " - ")
	return errors.New(fmt.Sprintf("%s - %s", ERR_NO_ENV, k))
}
