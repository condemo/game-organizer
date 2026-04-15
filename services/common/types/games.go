package types

import "time"

type IGDBCardResponse struct {
	IgdbID            int64     `json:"id"`
	Title             string    `json:"name"`
	InvolvedCompanies []int     `json:"involved_companies"`
	Cover             int       `json:"cover"`
	Platforms         []int     `json:"platforms"`
	ReleaseDate       time.Time `json:"first_release_date"`
}

type GameCard struct {
	Title       string
	Platforms   []string
	ReleaseDate time.Time
	Cover       string
	Developer   string
}

// TODO:
type IGDBGameResponse struct {
	IgdbID            int64     `json:"id"`
	Screenshots       []int     `json:"screenshots"`
	Title             string    `json:"name"`
	ReleaseDate       time.Time `json:"first_release_date"`
	Genres            []int     `json:"genres"`
	InvolvedCompanies []int     `json:"involved_companies"`
	Platforms         []int     `json:"platforms"`
	Rating            int       `json:"rating"`
	Url               string    `json:"url"`
}

type Game struct {
	ID          int64     `db:"id" json:"id"`
	IgdbID      int64     `db:"igdb_id" json:"igdb_id"`
	Title       string    `db:"title" json:"title"`
	Screenshot  string    `db:"screenshot" json:"screenshot"`
	ReleaseDate time.Time `db:"release_date" json:"release_date"`
	Genres      string    `db:"genres" json:"genres"`
	Developer   string    `db:"developer" json:"developer"`
	Publisher   string    `db:"publisher" json:"publisher"`
	Platforms   string    `db:"platforms" json:"platforms"`
	Rating      int       `db:"rating" json:"rating"`
	Url         string    `db:"url" json:"url"`
	Played      bool      `db:"player" json:"player"`
	Pending     bool      `db:"pending" json:"pending"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
