package fetch

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/condemo/game-organizer/services/common/config"
)

type twitchLoginRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func IgdbLogin() error {
	cl := http.Client{}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost,
		"https://id.twitch.tv/oauth2/token", nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("client_secret", os.Getenv("IGDB_CLIENT_SECRET"))
	q.Add("client_id", os.Getenv("IGDB_CLIENT_ID"))
	q.Add("grant_type", "client_credentials")
	req.URL.RawQuery = q.Encode()

	res, err := cl.Do(req)
	if err != nil {
		return err
	}

	var data twitchLoginRes
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	config.IGDBToken = data.AccessToken
	return nil
}
