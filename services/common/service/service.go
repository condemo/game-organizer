package service

import (
	"net/http"

	"github.com/condemo/game-organizer/services/common/store"
	"github.com/condemo/game-organizer/services/common/types"
)

type GameOrganizerService struct {
	store      store.Storage
	httpClient *http.Client
}

func NewGameOrganizerService(st store.Storage) *GameOrganizerService {
	hc := &http.Client{}
	return &GameOrganizerService{store: st, httpClient: hc}
}

func (s *GameOrganizerService) GetFetchGame(igdbID int64) (*types.Game, error) {
	return nil, nil
}

func (s *GameOrganizerService) GetOneGame(id int64) (*types.Game, error) {
	return nil, nil
}

func (s *GameOrganizerService) GetGames() ([]types.GameCard, error) {
	return nil, nil
}

func (s *GameOrganizerService) Search(q string) ([]types.GameCard, error) {
	return nil, nil
}

func (s *GameOrganizerService) CreateGame(g *types.Game) error {
	return nil
}

func (s *GameOrganizerService) UpdateGame(g *types.Game) error {
	return nil
}

func (s *GameOrganizerService) DeleteGame(id int64) error {
	return nil
}
