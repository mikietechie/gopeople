package config

import "os"

func GetEnvOrDef(key, _default string) string {
	env := os.Getenv(key)
	if len(env) != 0 {
		return env
	}
	return _default
}
