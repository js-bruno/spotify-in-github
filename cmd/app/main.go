package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/js-bruno/spotify-in-github/internal/services"
	"github.com/js-bruno/spotify-in-github/internal/util"
)

func main() {
	ctx := context.Background()
	env := util.LoadEnv()

	text, err := util.GenerateRandomString(100)
	if err != nil {
		return
	}
	hash := sha256.Sum256([]byte(text))
	bhash := hash[:]
	base64.RawURLEncoding.EncodeToString(bhash)

	// TODO: DONT NEED FOR NOW I THINK
	// response, err := services.GetClientCredentials(ctx, env.SpotifyClientId, env.SpotifyClientSecret)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for {
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
		songName := fmt.Sprint(current["item"].(map[string]any)["name"])
		songURL := current["item"].(map[string]any)["external_urls"].(map[string]any)["spotify"].(string)
		songArtists := current["item"].(map[string]any)["artists"].([]any)[0].(map[string]any)

		// websiteLayout := "by 𝑻𝒂𝒎𝒆 𝑰𝒎𝒑𝒂𝒍𝒂"
		// locationLayout := "i’m listening to 『𝑳𝒐𝒔𝒆𝒓』 on Spotify "
		songName = util.ConvertToFont(songName)
		artistName := util.ConvertToFont(songArtists["name"].(string))

		fmt.Println(songName)
		fmt.Println(IsPlaying)
		fmt.Println(deviceName)
		fmt.Println(songURL)
		fmt.Println(artistName)

		websiteLayout := fmt.Sprintf("by %s", artistName)
		locationLayout := fmt.Sprintf("🌴 i’m listening to 『%s』 on Spotify 🌴", songName)

		services.UpdateUserCompanyLocationWebsite(ctx, env.GithubTokenUser, locationLayout, websiteLayout, songURL)
		time.Sleep(30 * time.Second)
		fmt.Println("Waiting....")
	}
}
