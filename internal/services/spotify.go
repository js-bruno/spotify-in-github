package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

var TOKEN_URL = "https://accounts.spotify.com/api/token"
var CURRENTLY_PLAYING_URL = "https://api.spotify.com/v1/me/player"

var Client = &http.Client{Timeout: 2 * time.Second}

type AcessKeyResponse struct {
	AcessToken string `json:"access_token"`
	ExpiresIn  int    `json:"expires_in"`
}

func GetAccessToken(ctx context.Context, spotifyClientId, spotifyClientSecret, refreshToken string) (string, error) {
	payload := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	}
	req, _ := http.NewRequestWithContext(ctx, "POST", "https://accounts.spotify.com/api/token", strings.NewReader(payload.Encode()))
	auth := base64.StdEncoding.EncodeToString([]byte(spotifyClientId + ":" + spotifyClientSecret))

	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, _ := io.ReadAll(resp.Body)

	var result map[string]any
	json.Unmarshal(body, &result)

	return result["access_token"].(string), nil

}

func GetCurrentlyPlaying(ctx context.Context, accessToken string) (map[string]any, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", CURRENTLY_PLAYING_URL, nil)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var content map[string]any

	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return nil, err
	}

	return content, nil
}

func GetClientCredentials(ctx context.Context, clientID, clientSecret string) (AcessKeyResponse, error) {
	payload := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     clientID,
		"client_secret": clientSecret,
	}
	data := url.Values{}
	for key, value := range payload {
		data.Set(key, value)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", TOKEN_URL, strings.NewReader(data.Encode()))
	if err != nil {
		return AcessKeyResponse{}, err
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- REQUEST SENDING ---")
	fmt.Println(string(requestDump))
	fmt.Println("-----------------------")

	resp, err := Client.Do(req)
	if err != nil {
		return AcessKeyResponse{}, err
	}
	defer resp.Body.Close()

	var content AcessKeyResponse

	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return AcessKeyResponse{}, err
	}

	return content, nil
}
