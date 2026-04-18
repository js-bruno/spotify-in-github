package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"log"

	"github.com/js-bruno/spotify-in-github/internal/services"
	"github.com/js-bruno/spotify-in-github/internal/util"
)

var env = util.Env{}

func main() {
	env = util.LoadEnv()
	http.HandleFunc("/callback", HandleCallback)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/player", HandlePlayer)

	log.Printf("Waiting Callback in %s", env.CallbackPort)
	err := http.ListenAndServe(env.CallbackPort, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	scope := "user-read-playback-state"
	callback := fmt.Sprintf("http://127.0.0.1%s/callback", env.CallbackPort)
	auth := "https://accounts.spotify.com/authorize?" + url.Values{
		"response_type": {"code"},
		"client_id":     {env.SpotifyClientId},
		"scope":         {scope},
		"redirect_uri":  {callback},
	}.Encode()

	http.Redirect(w, r, auth, http.StatusFound)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	callback := fmt.Sprintf("http://127.0.0.1%s/callback", env.CallbackPort)
	data := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {callback},
	}

	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "Basic "+basicAuth())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var result map[string]any
	json.Unmarshal(body, &result)

	fmt.Println("ACCESS:", result["access_token"])
	fmt.Println("REFRESH:", result["refresh_token"])

	w.Write([]byte("COloque que o token no .env"))
}

func basicAuth() string {
	auth := env.SpotifyClientId + ":" + env.SpotifyClientSecret
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func HandlePlayer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	token, err := getAccessToken()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	response, err := services.GetCurrentlyPlaying(ctx, token)
	if err != nil {
		log.Println(err)
	}

	if len(response) == 0 {
		response = map[string]any{"status": "no-content"}
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getAccessToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", env.SpotifyRefreshToken)

	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token",
		strings.NewReader(data.Encode()))

	req.Header.Add("Authorization", "Basic "+basicAuth())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return result["access_token"].(string), nil
}
