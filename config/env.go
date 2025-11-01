package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Environment struct {
	PORT string

	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
	DB_USER     string
	DB_PASSWORD string
}

func InitEnvironment() Environment {
	// file .env or system environment
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file, so read the system env")
	}

	return read()
}

func InitEnvironmentWithFiles(filenames ...string) Environment {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatal("file not found")
	}

	return read()
}

func read() Environment {
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
