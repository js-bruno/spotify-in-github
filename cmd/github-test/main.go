package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/js-bruno/spotify-in-github/internal/util"
)

func main() {
	url := "https://api.github.com/user"

	jsonData := []byte(`{
		"bio": "Backend developer"
	}`)

	env := util.LoadEnv()
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	fmt.Println(env.GithubTokenUser)
	req.Header.Set("Authorization", "Bearer "+env.GithubTokenUser)
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
