package config

import (
	"fmt"
	"os"
)

func GetDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("AUTHOR_POSTGRES_USER"),
		os.Getenv("AUTHOR_POSTGRES_PASSWORD"),
		os.Getenv("AUTHOR_POSTGRES_HOST"),
		os.Getenv("AUTHOR_POSTGRES_PORT"),
		os.Getenv("AUTHOR_POSTGRES_DB"),
	)
}
