package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loadOnce sync.Once

func Config(key string) string {
	loadOnce.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Print("Error loading .env file")
		}
	})
	return os.Getenv(key)
}
