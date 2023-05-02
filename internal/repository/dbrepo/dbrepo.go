package dbrepo

import (
	"database/sql"

	"github.com/ZeCarlos88/bookings/internal/config"
	"github.com/ZeCarlos88/bookings/internal/repository"
)

type postgresDBrepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBrepo{
		App: a,
		DB:  conn,
	}
}
