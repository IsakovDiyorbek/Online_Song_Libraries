package models

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	ID          uuid.UUID `json:"id"`			
	GroupName   string    `json:"group_name"`        // Название группы
	SongName    string    `json:"song_name"`         // Название песни
	ReleaseDate string    `json:"release_date"` // Дата релиза
	Text        string    `json:"text"`         // Текст песни
	Link        string    `json:"link"`         // Ссылка на песню
	CreatedAt   time.Time `json:"created_at"`   // Дата создания записи
	UpdatedAt   time.Time `json:"updated_at"`   // Дата обновления записи
}


type AddSongRequest struct {
	GroupName   string    `json:"group_name"`        // Название группы
	SongName    string    `json:"song_name"`         // Название песни
	ReleaseDate string    `json:"release_date"` // Дата релиза
	Text        string    `json:"text"`         // Текст песни
	Link        string    `json:"link"`         // Ссылка на песню
}


type AddSongResponse struct {
    Success bool   `json:"success"`  // Успех операции
    Message string `json:"message"`  // Сообщение о результате
    SongID  string  `json:"id"`  // ID добавленной песни
}


type SongFilter struct{
	GroupName   string    `json:"group_name"`
	SongName    string    `json:"song_name"`   
	Page int
	PageSize int
}

