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
	ID          int64     `db:"id"`
	IgdbID      int64     `db:"igdb_id"`
	Title       string    `db:"title"`
	Screenshot  string    `db:"screenshot"`
	ReleaseDate time.Time `db:"release_date"`
	Genres      string    `db:"genres"`
	Developer   string    `db:"developer"`
	Publisher   string    `db:"publisher"`
	Platforms   string    `db:"platforms"`
	Rating      int       `db:"rating"`
	Url         string    `db:"url"`
	Played      bool      `db:"player"`
	Pending     bool      `db:"pending"`
	CreatedAt   time.Time `db:"created_at"`
}
