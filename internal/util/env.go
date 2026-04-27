package util

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() Env {
	godotenv.Load(".env.local")

	appDelaySeconds := os.Getenv("APP_DELAY_SECONDS")
	appDelaySecondsInt, err := strconv.Atoi(appDelaySeconds)

	if err != nil {
		log.Fatal(err.Error())
	}

	return Env{
		Enviroment:          os.Getenv("ENVIROMENT"),
		SpotifyClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyRefreshToken: os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		GithubTokenUser:     os.Getenv("GITHUB_API_TOKEN_USER"),
		CallbackPort:        os.Getenv("CALLBACK_PORT"),
		AppDelaySeconds:     appDelaySecondsInt,
	}
}

type Env struct {
	Enviroment          string
	SpotifyClientId     string
	SpotifyClientSecret string
	SpotifyRefreshToken string
	GithubTokenUser     string
	CallbackPort        string
	AppDelaySeconds     int
}
