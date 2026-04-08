package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := loadEnv()

	fmt.Println(env)
}

func loadEnv() Env {
	env := os.Getenv("ENVIRONMENT")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()

	return Env{
		Enviroment:    os.Getenv("ENVIROMENT"),
		SpotifyApiKey: os.Getenv("SPOTIFY_API_KEY"),
	}
}

type Env struct {
	Enviroment    string
	SpotifyApiKey string
}
