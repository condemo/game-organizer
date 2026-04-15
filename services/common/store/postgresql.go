package store

import (
	"fmt"
	"log"

	"github.com/condemo/game-organizer/services/common/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO:
func NewPosgresqlStore() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
		config.EnvConfig.Host, config.EnvConfig.User, config.EnvConfig.Pass,
		config.EnvConfig.Port, config.EnvConfig.Name)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
