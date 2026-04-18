package store

import (
	"context"

	"github.com/condemo/game-organizer/services/common/types"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetOneGame(id int64, game *types.Game) error
	GetGames() ([]types.Game, error)
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

func (s *Storage) GetOneGame(id int64, game *types.Game) error {
	q := `SELECT id, igdb_id, title, screenshot, cover, release_date, genres, developer, publisher,
	platforms, rating, url, played, pending, created_at FROM games WHERE id=$1;`
	if err := s.db.Get(game, q, id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetGamesPoltrait() ([]types.GamePoltrait, error) {
	var games []types.GamePoltrait
	q := `SELECT id, title, cover, played, pending from games;`
	if err := s.db.SelectContext(context.TODO(), &games, q); err != nil {
		return nil, err
	}
	return games, nil
}

func (s *Storage) CreateGame(g *types.Game) error {
	q := `INSERT INTO games
	(igdb_id,title,cover,screenshot,release_date,genres,developer,publisher,platforms,rating,
	url,played,pending)
	VALUES (:igdb_id,:title,:cover, :screenshot,:release_date,:genres,:developer,:publisher,:platforms,:rating,
	:url,:played,:pending) ON CONFLICT (title) DO NOTHING RETURNING *;`
	rows, err := s.db.NamedQueryContext(context.TODO(), q, g)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.StructScan(g); err != nil {
			return err
		}
	}

	if err := rows.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateGame(g *types.Game) error {
	q := `UPDATE games SET played=:played, pending=:pending WHERE id=:id RETURNING *;`
	rows, err := s.db.NamedQueryContext(context.TODO(), q, g)
	if err != nil {
		return err
	}

	if err := rows.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteGame(id int64) error {
	q := `DELETE FROM games WHERE id=$1`
	_, err := s.db.ExecContext(context.TODO(), q, id)
	if err != nil {
		return err
	}

	return nil
}
