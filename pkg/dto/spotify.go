package dto

type CurrentlyPlaying struct {
	Actions              Actions `json:"actions"`
	Context              Context `json:"context"`
	CurrentlyPlayingType string  `json:"currently_playing_type"`
	Device               Device  `json:"device"`
	IsPlaying            bool    `json:"is_playing"`
	Item                 Item    `json:"item"`
	ProgressMS           int     `json:"progress_ms"`
	RepeatState          string  `json:"repeat_state"`
	ShuffleState         bool    `json:"shuffle_state"`
	SmartShuffle         bool    `json:"smart_shuffle"`
	Timestamp            int64   `json:"timestamp"`
}

/* ---------- ACTIONS ---------- */

type Actions struct {
	Disallows Disallows `json:"disallows"`
}

type Disallows struct {
	Resuming bool `json:"resuming"`
}

/* ---------- CONTEXT ---------- */

type Context struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

/* ---------- DEVICE ---------- */

type Device struct {
	ID               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	SupportsVolume   bool   `json:"supports_volume"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
}

/* ---------- ITEM (TRACK) ---------- */

type Item struct {
	Album            Album        `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMS       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIDs      ExternalIDs  `json:"external_ids"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsLocal          bool         `json:"is_local"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       *string      `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
}

/* ---------- ALBUM ---------- */

type Album struct {
	AlbumType        string       `json:"album_type"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	Images           []Image      `json:"images"`
	Name             string       `json:"name"`
	ReleaseDate      string       `json:"release_date"`
	ReleasePrecision string       `json:"release_date_precision"`
	TotalTracks      int          `json:"total_tracks"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
}

/* ---------- SHARED TYPES ---------- */

type Artist struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type ExternalIDs struct {
	ISRC string `json:"isrc"`
}

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
