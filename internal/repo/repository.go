package repo

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{}
}
