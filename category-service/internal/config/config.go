package config

import (
	"fmt"
	"os"
)

func GetDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("CATEGORY_POSTGRES_USER"),
		os.Getenv("CATEGORY_POSTGRES_PASSWORD"),
		os.Getenv("CATEGORY_POSTGRES_HOST"),
		os.Getenv("CATEGORY_POSTGRES_PORT"),
		os.Getenv("CATEGORY_POSTGRES_DB"),
	)
}
