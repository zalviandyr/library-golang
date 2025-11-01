package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Environment struct {
	PORT              string
	DB_HOST           string
	DB_PORT           string
	DB_DATABASE       string
	DB_USER           string
	DB_PASSWORD       string
	STORAGE_DRIVER    string
	MINIO_ENDPOINT    string
	MINIO_SERVER_HOST string
	MINIO_BUCKET      string
	MINIO_ACCESS_KEY  string
	MINIO_SECRET_KEY  string
}

func InitEnvironment() Environment {
	// file .env or system environment
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file, so read the system env")
	}

	envContent, err := godotenv.Read()
	if err != nil {
		fmt.Println("error reading .env, so read the system env")

		// read system environment
		envContent = map[string]string{}
		for _, env := range os.Environ() {
			data := strings.Split(env, "=")
			key := data[0]
			value := data[1]

			envContent[key] = value
		}
	}

	env := Environment{}
	jsonBody, _ := json.Marshal(envContent)
	json.Unmarshal(jsonBody, &env)

	return env
}
