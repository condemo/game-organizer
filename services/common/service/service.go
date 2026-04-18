package service

import (
	"github.com/condemo/game-organizer/services/common/fetch"
	"github.com/condemo/game-organizer/services/common/store"
	"github.com/condemo/game-organizer/services/common/types"
)

type GameOrganizerService struct {
	st          store.Storage
	gameFetcher *fetch.GameFetcher
}

func NewGameOrganizerService(st store.Storage) *GameOrganizerService {
	return &GameOrganizerService{st: st, gameFetcher: fetch.NewGameFetcher()}
}

func (s *GameOrganizerService) GetFetchGame(igdbID int64) (*types.Game, error) {
	// get raw game with gameFetcher
	// convertion to types.Game
	// return to handler
	return nil, nil
}

func (s *GameOrganizerService) GetOneGame(id int64) (*types.Game, error) {
	var game *types.Game

	if err := s.st.GetOneGame(id, game); err != nil {
		return nil, err
	}

	return game, nil
}

// TODO:
func (s *GameOrganizerService) GetGames() ([]types.GamePoltrait, error) {
	games, err := s.st.GetGamesPoltrait()
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (s *GameOrganizerService) Search(q string) ([]types.GameCard, error) {
	// get raw game with gameFetcher
	// convertion to []types.GameCard
	// return to handler
	return nil, nil
}

func (s *GameOrganizerService) CreateGame(g *types.Game) error {
	if err := s.st.CreateGame(g); err != nil {
		return err
	}
	return nil
}

// TODO:
func (s *GameOrganizerService) UpdateGame(g *types.Game) error {
	return nil
}

func (s *GameOrganizerService) DeleteGame(id int64) error {
	if err := s.st.DeleteGame(id); err != nil {
		return err
	}
	return nil
}
