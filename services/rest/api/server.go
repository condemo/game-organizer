package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/condemo/game-organizer/services/common/service"
	"github.com/condemo/game-organizer/services/common/store"
	"github.com/condemo/game-organizer/services/rest/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}

func (s *ApiServer) Run() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://192.168.3.*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowCredentials: false,
		MaxAge:           300,
	}), middleware.Logger)

	// Database
	pgDB := store.NewPosgresqlStore()
	store := store.NewStorage(pgDB)

	// Service
	service := service.NewGameOrganizerService(*store)

	// Handlers
	gameHandler := handlers.NewGameHandler(service)
	r.Mount("/games", gameHandler.RegisterRoutes())

	server := http.Server{
		Addr:         s.addr,
		Handler:      r,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	log.Printf("server running on port %s", s.addr)

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-sigC

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("bad shutdown: ", err)
	}
}
