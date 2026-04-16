package store

import (
	"context"

	"github.com/condemo/game-organizer/services/common/types"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetOneGame(id int64, game *types.Game) error
	GetGames() ([]types.Game, error)
	Search(q string) []types.GameCard
	CreateGame(g *types.Game) error
	UpdateGame(g *types.Game) error
	DeleteGame(id int64) error
}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// TODO:
func (s *Storage) GetOneGame(id int64, game *types.Game) error {
	q := `SELECT id, igdb_id, title, screenshot, release_date, genres, developer, publisher,
	platforms, rating, url, played, pending, created_at FROM games WHERE id=$1`
	if err := s.db.Get(game, q, id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetGames() ([]types.Game, error) {
	return nil, nil
}

func (s *Storage) Search(q string) []types.GameCard {
	return nil
}

func (s *Storage) CreateGame(g *types.Game) error {
	q := `INSERT INTO games
	(igdb_id,title,screenshot,release_date,genres,developer,publisher,platforms,rating,
	url,played,pending)
	VALUES (:igdb_id,:title,:screenshot,:release_date,:genres,:developer,:publisher,:platforms,:rating,
	:url,:played,:pending) ON CONFLICT (title) DO NOTHING RETURNING *`
	rows, err := s.db.NamedQueryContext(context.TODO(), q, g)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.StructScan(g); err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) UpdateGame(g *types.Game) error {
	return nil
}

func (s *Storage) DeleteGame(id int64) error {
	return nil
}
