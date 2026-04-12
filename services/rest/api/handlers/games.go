package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

type GameHandler struct{}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

func (h *GameHandler) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.HandleFunc("/test", MakeHandler(h.test))

	return r
}

func (h *GameHandler) test(w http.ResponseWriter, r *http.Request) error {
	JsonResponse(w, http.StatusOK, "working")
	return nil
}
