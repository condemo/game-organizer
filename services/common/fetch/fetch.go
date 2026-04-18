package fetch

import (
	"net/http"

	"github.com/condemo/game-organizer/services/common/types"
)

type GameFetcher struct {
	client *http.Client
}

func NewGameFetcher() *GameFetcher {
	return &GameFetcher{
		client: &http.Client{},
	}
}

// TODO:
func (gf *GameFetcher) Search(q string) ([]types.IGDBCardResponse, error) {
	return nil, nil
}

// TODO:
func (gf *GameFetcher) FetchOneGame(id int64) (*types.IGDBGameResponse, error) {
	return nil, nil
}
