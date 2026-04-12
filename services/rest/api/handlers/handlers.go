package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/condemo/game-organizer/services/rest/api/errs"
)

type customHandler func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(f customHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			switch err := err.(type) {
			case errs.ApiError:
				JsonResponse(w, err.Status, err.Message)
			default:
				JsonResponse(w, http.StatusInternalServerError, errs.InternalServerError)
			}
			slog.Error(err.Error(), "path", r.URL.Path)
		}
	}
}

func JsonResponse(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		JsonResponse(w, http.StatusInternalServerError, errs.InternalServerError)
	}
}
