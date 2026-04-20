package service

import (
	"github.com/condemo/game-organizer/services/common/service/fetch"
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
	gameRes, err := s.gameFetcher.FetchOneGame(igdbID)
	if err != nil {
		return nil, err
	}
	// convertion to types.Game
	game := gameRes.GetDBStruct()
	// return to handler
	return &game, nil
}

func (s *GameOrganizerService) GetOneGame(id int64) (*types.Game, error) {
	var game *types.Game

	if err := s.st.GetOneGame(id, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *GameOrganizerService) GetGames() ([]types.GamePoltrait, error) {
	games, err := s.st.GetGamesPoltrait()
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (s *GameOrganizerService) Search(q string) ([]types.GameCard, error) {
	// get raw game with gameFetcher
	gameRes, err := s.gameFetcher.Search(q)
	if err != nil {
		return nil, err
	}
	// convertion to []types.GameCard
	var gamesCard []types.GameCard
	for _, g := range gameRes {
		gamesCard = append(gamesCard, g.GetDBStruct())
	}
	// return to handler
	return gamesCard, nil
}

func (s *GameOrganizerService) CreateGame(g *types.Game) error {
	if err := s.st.CreateGame(g); err != nil {
		return err
	}
	return nil
}

func (s *GameOrganizerService) UpdateGame(g *types.Game) error {
	if err := s.st.UpdateGame(g); err != nil {
		return err
	}
	return nil
}

func (s *GameOrganizerService) DeleteGame(id int64) error {
	if err := s.st.DeleteGame(id); err != nil {
		return err
	}
	return nil
}
