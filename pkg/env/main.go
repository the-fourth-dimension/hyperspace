package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const PORT = iota

var keys = []string{"PORT"}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprint("Error loading .env file: ", err))
	}
	for idx := range keys {
		key := keys[idx]
		_, isPresent := os.LookupEnv(key)
		if !isPresent {
			panic(fmt.Sprintf("missing environment variable: %s", key))
		}

	}
}

func GetEnv(key int) string {
	return os.Getenv(keys[key])
}
