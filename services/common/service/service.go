package service

import "github.com/condemo/game-organizer/services/common/store"

type GameOrganizerService struct {
	store store.Storage
}

func NewGameOrganizerService(st store.Storage) *GameOrganizerService {
	return &GameOrganizerService{store: st}
}
