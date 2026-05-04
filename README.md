![a](md-gifs/vine4.gif)
# spotify-github-status

> Automatically updates your GitHub profile with the song currently playing on Spotify.

![preview](https://img.shields.io/badge/status-active-brightgreen)
![license](https://img.shields.io/badge/license-MIT-blue)
![go](https://img.shields.io/badge/go-%3E%3D1.21-00ADD8)

---

## ✨ How it works

The project polls the Spotify API at regular intervals and, whenever it detects a song playing, updates your GitHub profile with the track name and artist — all running automatically in the background.

```
Spotify API → get song in the player → GitHub API → updates profile
```

**Example of generated status:**

> 🌴 i’m listening to 『𝑨1 - 𝑰𝒕'𝒔 𝒋𝒖𝒔𝒕 𝒂 𝒃𝒖𝒓𝒏𝒊𝒏𝒈 𝒎𝒆𝒎𝒐𝒓𝒚』 on Spotify 🌴
> by The Caretaker
---

## installation

### pre-requisites

- Go >= 1.21
- A [Spotify for Developers](https://developer.spotify.com/dashboard) account
- GitHub Personal Access Token with `user` scope

### steps

```bash
# 1. Clone the repository
git clone https://github.com/your-username/spotify-github-status.git
cd spotify-github-status

# 2. Build the project
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o spotify-in-github cmd/app/main.go

# 3. Set up environment variables
cp .env.example .env.local
```

---

## configuration

Fill in the `.env.local` file with your credentials:

```env
# Spotify
SPOTIFY_CLIENT_ID=your_client_id
SPOTIFY_CLIENT_SECRET=your_client_secret
SPOTIFY_REFRESH_TOKEN=your_refresh_token

# GitHub
GITHUB_TOKEN=your_personal_access_token
GITHUB_USERNAME=your_username
```

### getting spotify credentials

1. Go to the [Spotify Developer Dashboard](https://developer.spotify.com/dashboard) and create an app
2. Copy the **Client ID** and **Client Secret**
3. Add `http://localhost:3000/callback` as a Redirect URI
4. Run the auth flow to get your **Refresh Token**:

```bash
go run ./cmd/app/main.go
```

### getting the github token

1. Go to **GitHub → Settings → Developer settings → Personal access tokens**
2. Create a token with the `user` scope enabled
3. Paste the token into your `.env`

---

## usage

```bash
# Run normally
./spotify-github-status

# Or run directly with Go
go run ./cmd/...
```

The process will keep running and update your GitHub profile automatically. When no song is playing, the status is cleared.

---

## contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

1. Fork the project
2. Create a branch: `git checkout -b feat/my-feature`
3. Commit your changes: `git commit -m 'feat: add my feature'`
4. Push to the branch: `git push origin feat/my-feature`
5. Open a Pull Request

---

## license

Distributed under the MIT License. See [`LICENSE`](LICENSE) for more information.

---

<p align="center">Made with 💊 and 🎵</p>
![a](md-gifs/vine4.gif)
