package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/js-bruno/spotify-in-github/internal/services"
	"github.com/js-bruno/spotify-in-github/internal/util"
)

func main() {
	ctx := context.Background()
	env := util.LoadEnv()

	text, err := util.GenerateRandomString(100)
	hash := sha256.Sum256([]byte(text))
	bhash := hash[:]
	base64.RawURLEncoding.EncodeToString(bhash)

	// TODO: DONT NEED FOR NOW I THINK
	// response, err := services.GetClientCredentials(ctx, env.SpotifyClientId, env.SpotifyClientSecret)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	accessToken, err := services.GetAccessToken(ctx, env.SpotifyClientId, env.SpotifyClientSecret, env.SpotifyRefreshToken)
	if err != nil {
		log.Fatal(err)
	}

	current, err := services.GetCurrentlyPlaying(ctx, accessToken)
	if err != nil {
		log.Fatal(err)
	}

	IsPlaying := current["is_playing"]
	deviceName := current["device"].(map[string]any)["type"]
	songName := current["item"].(map[string]any)["name"]

	fmt.Println(IsPlaying)
	fmt.Println(deviceName)
	fmt.Println(songName)

}
