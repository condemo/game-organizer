package fetch

import (
	"strings"
	"time"

	"github.com/condemo/game-organizer/services/common/types"
)

type igdbCardResponse struct {
	IgdbID            int64  `json:"id"`
	Title             string `json:"name"`
	InvolvedCompanies []struct {
		Developer bool `json:"developer"`
		Publisher bool `json:"publisher"`
		Company   struct {
			Name string `json:"name"`
		} `json:"company"`
	} `json:"involved_companies"`
	Cover struct {
		Url string `json:"url"`
	} `json:"cover"`
	Platforms []struct {
		Name string `json:"name"`
	} `json:"platforms"`
	ReleaseDate int `json:"first_release_date"`
}

func (c igdbCardResponse) GetDBStruct() types.GameCard {
	gc := types.GameCard{
		ID:    c.IgdbID,
		Title: c.Title,
	}

	// developer
	for _, d := range c.InvolvedCompanies {
		if d.Developer {
			gc.Developer = d.Company.Name
		}
	}

	// platforms
	var pl []string
	for _, p := range c.Platforms {
		pl = append(pl, p.Name)
	}
	gc.Platforms = strings.Join(pl, ",")

	// ReleaseDate
	gc.ReleaseDate = time.Unix(int64(c.ReleaseDate), 0)

	// Cover
	gc.Cover = c.Cover.Url[2:]

	return gc
}

type igdbGameResponse struct {
	IgdbID      int64 `json:"id"`
	Screenshots []struct {
		Url string `json:"url"`
	} `json:"screenshots"`
	Title       string `json:"name"`
	ReleaseDate int    `json:"first_release_date"`
	Genres      []struct {
		Name string `json:"name"`
	} `json:"genres"`
	Cover struct {
		Url string `json:"url"`
	} `json:"cover"`
	InvolvedCompanies []struct {
		Developer bool `json:"developer"`
		Publisher bool `json:"publisher"`
		Company   struct {
			Name string `json:"name"`
		} `json:"company"`
	} `json:"involved_companies"`
	Platforms []struct {
		Name string `json:"name"`
	} `json:"platforms"`
	Rating float32 `json:"rating"`
	Url    string  `json:"url"`
}

func (c igdbGameResponse) GetDBStruct() types.Game {
	game := types.Game{
		IgdbID: c.IgdbID,
		Title:  c.Title,
		Url:    c.Url,
		Rating: int(c.Rating),
	}

	// Screenshots
	game.Screenshot = c.Screenshots[0].Url[2:]

	// Cover
	game.Cover = c.Cover.Url[2:]

	// ReleaseDate
	game.ReleaseDate = time.Unix(int64(c.ReleaseDate), 0)

	// Genres
	var genres []string
	for _, g := range c.Genres {
		genres = append(genres, g.Name)
	}
	game.Genres = strings.Join(genres, ",")

	// Developer | Publisher
	for _, c := range c.InvolvedCompanies {
		if c.Developer {
			game.Developer = c.Company.Name
		}
		if c.Publisher {
			game.Publisher = c.Company.Name
		}
	}

	// Platforms
	var pl []string
	for _, p := range c.Platforms {
		pl = append(pl, p.Name)
	}
	game.Platforms = strings.Join(pl, ",")

	return game
}
