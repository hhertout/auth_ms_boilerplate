package repository

import (
	"auth_ms/cmd/database"
	"database/sql"
)

type Repository struct {
	dbPool *sql.DB
}

func NewRepository(customSource *sql.DB) (*Repository, error) {
	if customSource != nil {
		return &Repository{
			customSource,
		}, nil
	} else {
		dbService, err := database.Connect()
		if err != nil {
			return nil, err
		}

		return &Repository{
			dbService.DbPool,
		}, nil
	}
}
