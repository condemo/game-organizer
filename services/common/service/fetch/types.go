package fetch

import "github.com/condemo/game-organizer/services/common/types"

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
	Cover     int `json:"cover"`
	Platforms []struct {
		Name string `json:"name"`
	} `json:"platforms"`
	ReleaseDate int `json:"first_release_date"`
}

// TODO:
func (c igdbCardResponse) GetDBStruct() types.GameCard {
	return types.GameCard{}
}

type igdbGameResponse struct {
	IgdbID            int64  `json:"id"`
	Screenshots       []int  `json:"screenshots"`
	Title             string `json:"name"`
	ReleaseDate       int    `json:"first_release_date"`
	Genres            []int  `json:"genres"`
	Cover             int    `json:"cover"`
	InvolvedCompanies []struct {
		Developer bool `json:"developer"`
		Publisher bool `json:"publisher"`
		Company   struct {
			Name string `json:"name"`
		} `json:"company"`
	} `json:"involved_companies"`
	Platforms []int  `json:"platforms"`
	Rating    int    `json:"rating"`
	Url       string `json:"url"`
}

// TODO:
func (c igdbGameResponse) GetDbStruct() types.Game {
	return types.Game{}
}
