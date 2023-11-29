package dbrepo

import (
	"database/sql"

	"github.com/tsawler/bookings-app/internal/config"
	"github.com/tsawler/bookings-app/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, appConfig *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: appConfig,
		DB:  conn,
	}
}
