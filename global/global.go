package global

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)


func main() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	
	var SPOTIFY_ID = os.Getenv("SPOTIFY_ID")
	var SPOTIFY_SECRET = os.Getenv("SPOTIFY_SECRET")
	const RedirectURI = "http://localhost:8080/callback"

	
	var (
		Auth  = spotifyauth.New(
			spotifyauth.WithRedirectURL(RedirectURI),
			spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserModifyPlaybackState),
			spotifyauth.WithClientID(SPOTIFY_ID),
			spotifyauth.WithClientSecret(SPOTIFY_SECRET),
			)
		Ch    = make(chan *spotify.Client)
		State = "abc123"
	)
}