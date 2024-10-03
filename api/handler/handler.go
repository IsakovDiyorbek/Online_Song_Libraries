package handler

import (
	"database/sql"

	"github.com/Online_Song_Libraries/storage/postgres"
)



type Handler struct{
	db *sql.DB
	song *postgres.SongLibrary
}

func NewHandler(db *sql.DB) *Handler{
	return &Handler{
		db: db,
		song: postgres.NewSongLibrary(db),
	}
}