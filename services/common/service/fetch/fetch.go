package fetch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/condemo/game-organizer/services/common/config"
)

type GameFetcher struct {
	client  *http.Client
	gameUrl string
}

func NewGameFetcher() *GameFetcher {
	return &GameFetcher{
		client:  &http.Client{},
		gameUrl: "https://api.igdb.com/v4/games",
	}
}

// TODO:
func (gf *GameFetcher) Search(q string) ([]igdbCardResponse, error) {
	fields := `fields id,name,cover.url,platforms,first_release_date,
	involved_companies.developer, involved_companies.publisher,
	involved_companies.company.name,platforms.name;`
	bodyStr := fmt.Sprintf(
		"search \"%s\";%s limit 5;", q, fields)
	req, err := http.NewRequestWithContext(context.TODO(),
		http.MethodPost,
		gf.gameUrl,
		strings.NewReader(bodyStr),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Client-ID", os.Getenv("IGDB_CLIENT_ID"))
	req.Header.Add("Authorization", "Bearer "+config.IGDBToken)

	res, err := gf.client.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO: crear un switch que contemple los distintos status codes para devolver errores
	// mas precisos y en caso de un 401 intente logearse una vez
	if res.StatusCode == http.StatusOK {
		var cardRes []igdbCardResponse
		if err := json.NewDecoder(res.Body).Decode(&cardRes); err != nil {
			return nil, err
		}
		return cardRes, nil
	} else {
		errStr := fmt.Sprintf("igbd error status code: %s", res.Status)
		return nil, errors.New(errStr)
	}
}

// TODO:
func (gf *GameFetcher) FetchOneGame(id int64) (*igdbGameResponse, error) {
	fields := `fields id,name,cover.url,rating,url,
	screenshots.url,genres.name,involved_companies.developer,involved_companies.publisher,
	involved_companies.company.name,platforms.name;`
	bodyStr := fmt.Sprintf("%s where id = %d;", fields, id)

	req, err := http.NewRequestWithContext(context.TODO(),
		http.MethodPost,
		gf.gameUrl,
		strings.NewReader(bodyStr),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+config.IGDBToken)
	req.Header.Add("Client-ID", os.Getenv("IGDB_CLIENT_ID"))

	res, err := gf.client.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO: crear un switch que contemple los distintos status codes para devolver errores
	// mas precisos y en caso de un 401 intente logearse una vez
	if res.StatusCode == http.StatusOK {
		var gameRes []igdbGameResponse
		if err := json.NewDecoder(res.Body).Decode(&gameRes); err != nil {
			return nil, err
		}
		if len(gameRes) > 0 {
			return &gameRes[0], nil
		} else {
			return nil, errors.New("empty game array response")
		}
	} else {
		fmt.Printf("something went wrong: %s\n", res.Status)
	}

	return nil, nil
}
