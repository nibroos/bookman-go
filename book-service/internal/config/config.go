package config

import (
	"fmt"
	"os"
)

func GetDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("BOOK_POSTGRES_USER"),
		os.Getenv("BOOK_POSTGRES_PASSWORD"),
		os.Getenv("BOOK_POSTGRES_HOST"),
		os.Getenv("BOOK_POSTGRES_PORT"),
		os.Getenv("BOOK_POSTGRES_DB"),
	)
}
