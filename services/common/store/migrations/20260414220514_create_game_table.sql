-- +goose Up
CREATE TABLE IF NOT EXISTS games (
  id SERIAL PRIMARY KEY,
  igdb_id INTEGER NOT NULL,
  title VARCHAR(80) NOT NULL,
  screenshot VARCHAR,
  release_date DATE,
  genres VARCHAR,
  developer VARCHAR(80),
  publisher VARCHAR(80),
  platforms VARCHAR,
  rating SMALLINT,
  url VARCHAR(200),
  played BOOLEAN DEFAULT false,
  pending BOOLEAN DEFAULT false,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
  UNIQUE (title)
);

-- +goose Down
DROP TABLE game;
