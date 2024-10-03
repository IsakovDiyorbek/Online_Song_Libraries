package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Online_Song_Libraries/models"
	"github.com/google/uuid"
)

type SongLibrary struct {
	db *sql.DB
}

func NewSongLibrary(db *sql.DB) *SongLibrary {
	return &SongLibrary{
		db: db,
	}
}

func (s *SongLibrary) AddSong(req *models.AddSongRequest) (*models.AddSongResponse, error) {
	log.Printf("Adding new song: GroupName=%s, SongName=%s, ReleaseDate=%s, Link=%s", req.GroupName, req.SongName, req.ReleaseDate, req.Link)

	query := `INSERT INTO songs(id, group_name, song_name, release_date, text, link) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	Id := uuid.NewString()
	var songID string
	err := s.db.QueryRow(query, Id, req.GroupName, req.SongName, req.ReleaseDate, req.Text, req.Link).Scan(&songID)
	if err != nil {
		log.Println("Failed to add song: %v", err)
		return &models.AddSongResponse{
			Success: false,
			Message: "Failed to add song",
		}, fmt.Errorf("failed to add song: %v", err)
	}

	return &models.AddSongResponse{
		Success: true,
		Message: "Song added successfully",
		SongID:  songID,
	}, nil
}



func (s *SongLibrary) GetSongs(filter models.SongFilter) ([]models.Song, error) {
    query := `SELECT id, group_name, song_name, release_date, link 
              FROM songs 
              WHERE ($1::text IS NULL OR group_name ILIKE $1) 
              AND ($2::text IS NULL OR song_name ILIKE $2) 
              LIMIT $3 OFFSET $4`

    offset := (filter.Page - 1) * filter.PageSize
    rows, err := s.db.Query(query, filter.GroupName, filter.SongName, filter.PageSize, offset)
    if err != nil {
		log.Println("Error executing query: %v", err)
        return nil, err
    }
    defer rows.Close()

    var songs []models.Song
    for rows.Next() {
        var song models.Song
        if err := rows.Scan(&song.ID, &song.GroupName, &song.SongName, &song.ReleaseDate, &song.Link); err != nil {
			log.Println("Error scanning row: %v", err)
            return nil, err
        }
        songs = append(songs, song)
    }

    return songs, nil
}