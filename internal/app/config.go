package app

import "os"

type Config struct {
	DatabaseURL string
}

// New fetches configuration values and builds a new app config from them
func NewConfig() Config {
	var (
		DatabaseURL = mustGetString("DATABASE_URL")
	)

	return Config{
		DatabaseURL: DatabaseURL,
	}
}

func mustGetString(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	panic("Value for key '" + key + "' must be provided.")
}
