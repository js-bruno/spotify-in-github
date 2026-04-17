package util

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() Env {
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
		Enviroment:          os.Getenv("ENVIROMENT"),
		SpotifyClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyRefreshToken: os.Getenv("SPOTIFY_REFRESH_TOKEN"),
		GithubToken:         os.Getenv("GITHUB_TOKEN"),
		CallbackPort:        os.Getenv("CALLBACK_PORT"),
	}
}

type Env struct {
	Enviroment          string
	SpotifyClientId     string
	SpotifyClientSecret string
	SpotifyRefreshToken string
	GithubToken         string
	CallbackPort        string
}
