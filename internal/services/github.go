package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var USER_API_URL = "https://api.github.com/user"

func UpdateUserCompanyLocationWebsite(ctx context.Context, githubToken, location, website, twitterUser string) error {
	// jsonData := []byte(`{
	// 	"company": "Backend developer"
	// }`)
	data := map[string]string{
		"blog":             website,
		"location":         location,
		"twitter_username": twitterUser,
	}

	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("PATCH", USER_API_URL, bytes.NewBuffer(jsonData))
	req.Header.Add("Authorization", "Bearer "+githubToken)
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil

	}
	return errors.New("Update Fails: " + resp.Status)
}
