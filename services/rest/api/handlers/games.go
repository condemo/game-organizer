package handlers

import (
	"net/http"

	"github.com/condemo/game-organizer/services/common/service"
	"github.com/go-chi/chi"
)

type GameHandler struct {
	service *service.GameOrganizerService
}

func NewGameHandler(s *service.GameOrganizerService) *GameHandler {
	return &GameHandler{service: s}
}

func (h *GameHandler) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", MakeHandler(h.getGames))
	r.Get("/{id}", MakeHandler(h.getOneGame))
	r.Get("/search", MakeHandler(h.search))
	r.Put("/", MakeHandler(h.createGame))
	r.Delete("/", MakeHandler(h.deleteGame))
	r.Put("/", MakeHandler(h.updateGame))

	return r
}

// TODO:
func (h *GameHandler) getOneGame(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// TODO:
func (h *GameHandler) getGames(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// TODO:
func (h *GameHandler) search(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// TODO:
func (h *GameHandler) createGame(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// TODO:
func (h *GameHandler) updateGame(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// TODO:
func (h *GameHandler) deleteGame(w http.ResponseWriter, r *http.Request) error {
	return nil
}
