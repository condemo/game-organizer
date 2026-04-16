package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/condemo/game-organizer/services/common/service"
	"github.com/condemo/game-organizer/services/common/types"
	"github.com/condemo/game-organizer/services/rest/api/errs"
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
	r.Get("/fetch/{id}", MakeHandler(h.getFetchGame))
	r.Get("/{id}", MakeHandler(h.getOneGame))
	r.Get("/", MakeHandler(h.getGames))
	r.Get("/search", MakeHandler(h.search))
	r.Put("/", MakeHandler(h.createGame))
	r.Put("/", MakeHandler(h.updateGame))
	r.Delete("/{id}", MakeHandler(h.deleteGame))

	return r
}

func (h *GameHandler) getFetchGame(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return errs.NewApiError(err, http.StatusBadRequest, "game id must be a number")
	}

	game, err := h.service.GetFetchGame(id)
	if err != nil {
		return err
	}

	JsonResponse(w, http.StatusOK, game)

	return nil
}

func (h *GameHandler) getOneGame(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return errs.NewApiError(err, http.StatusBadRequest, "game id must be a number")
	}

	game, err := h.service.GetOneGame(id)
	if err != nil {
		return err
	}

	JsonResponse(w, http.StatusOK, game)

	return nil
}

func (h *GameHandler) getGames(w http.ResponseWriter, r *http.Request) error {
	games, err := h.service.GetGames()
	if err != nil {
		return err
	}

	JsonResponse(w, http.StatusOK, games)
	return nil
}

func (h *GameHandler) search(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("q")
	res, err := h.service.Search(q)
	if err != nil {
		return err
	}

	JsonResponse(w, http.StatusOK, res)
	return nil
}

func (h *GameHandler) createGame(w http.ResponseWriter, r *http.Request) error {
	var game types.Game

	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		return errs.NewApiError(err, http.StatusBadRequest, "invalid json format")
	}

	if err := h.service.CreateGame(&game); err != nil {
		return err
	}

	JsonResponse(w, http.StatusCreated, game)
	return nil
}

func (h *GameHandler) updateGame(w http.ResponseWriter, r *http.Request) error {
	var game types.Game

	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		return errs.NewApiError(err, http.StatusBadRequest, "invalid json format")
	}

	if err := h.service.UpdateGame(&game); err != nil {
		return err
	}

	JsonResponse(w, http.StatusCreated, game)
	return nil
}

func (h *GameHandler) deleteGame(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return errs.NewApiError(err, http.StatusBadRequest, "game id must be a number")
	}

	if err := h.service.DeleteGame(id); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}
