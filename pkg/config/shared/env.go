package shared

import (
	"os"
	"strings"

	"github.com/s3rj1k/ninit/pkg/validate"
)

// LookupEnvValue is a wrapper around `os.LookupEnv` and `validate.EnvironName`.
func LookupEnvValue(env string) (val string, ok bool, err error) {
	err = validate.EnvironName(env)
	if err != nil {
		return "", false, err //nolint: wrapcheck // error string formed in external package is styled correctly
	}

	val, ok = os.LookupEnv(env)
	if !ok {
		return "", false, nil
	}

	if strings.TrimSpace(val) == "" {
		return "", false, nil
	}

	return val, true, nil
}
