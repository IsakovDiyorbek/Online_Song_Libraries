package models

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	ID          string `json:"id"`			
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


type UpdateSongRequest struct {
	ID          uuid.UUID `json:"id"`		
    GroupName   string `json:"group_name,omitempty"`
    SongName    string `json:"song_name,omitempty"`
    ReleaseDate string `json:"release_date,omitempty"`
    Text        string `json:"text,omitempty"`
    Link        string `json:"link,omitempty"`
}


type GetAllSongsRequest struct {
    GroupName string `json:"group_name"`
    SongName  string `json:"song_name"`
    Text      string `json:"text"`
    Limit     int    `json:"limit"`
    Offset    int    `json:"offset"`
}



type VerseResponse struct {
	Id   string  `json:"id"`
	VersNum  int    `json:"verseNum"`
	Text     string `json:"text"`
}
