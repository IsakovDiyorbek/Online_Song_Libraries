package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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

func (s *SongLibrary) GetSongs(filter *models.SongFilter) (*[]models.Song, error) {
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

	return &songs, nil
}
func (s *SongLibrary) GetAll(req *models.GetAllSongsRequest) ([]*models.Song, error) {
    query := `SELECT id, group_name, song_name, release_date, text, link FROM songs WHERE 1=1`
    var args []interface{}
    var filters []string

    if req.GroupName != "" {
        filters = append(filters, fmt.Sprintf("group_name ILIKE $%d", len(args)+1))
        args = append(args, "%"+req.GroupName+"%")
    }

    if req.SongName != "" {
        filters = append(filters, fmt.Sprintf("song_name ILIKE $%d", len(args)+1))
        args = append(args, "%"+req.SongName+"%")
    }

    if req.Text != "" {
        filters = append(filters, fmt.Sprintf("text ILIKE $%d", len(args)+1))
        args = append(args, "%"+req.Text+"%")
    }

    if len(filters) > 0 {
        query += " AND " + strings.Join(filters, " AND ")
    }

    query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
    args = append(args, req.Limit, req.Offset)


    rows, err := s.db.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("failed to get songs: %v", err)
    }
    defer rows.Close()

    var songs []*models.Song
    for rows.Next() {
        var song models.Song
        if err := rows.Scan(&song.ID, &song.GroupName, &song.SongName, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
            return nil, fmt.Errorf("failed to scan song: %v", err)
        }
        songs = append(songs, &song)
    }

    return songs, nil
}

func (s *SongLibrary) DeleteSong(songID string) error {
	query := `DELETE FROM songs WHERE id = $1`
	_, err := s.db.Exec(query, songID)
	if err != nil {
		log.Println("Error executing query: %v", err)
		return err
	}
	return nil
}

func (s *SongLibrary) UpdateSong(songID string, req models.UpdateSongRequest) error {
	query := `UPDATE songs SET`
	params := []interface{}{}
	paramCount := 1

	if req.GroupName != "" {
		query += fmt.Sprintf(" group_name = $%d,", paramCount)
		params = append(params, req.GroupName)
		paramCount++
	}
	if req.SongName != "" {
		query += fmt.Sprintf(" song_name = $%d,", paramCount)
		params = append(params, req.SongName)
		paramCount++
	}
	if req.ReleaseDate != "" {
		query += fmt.Sprintf(" release_date = $%d,", paramCount)
		params = append(params, req.ReleaseDate)
		paramCount++
	}
	if req.Text != "" {
		query += fmt.Sprintf(" text = $%d,", paramCount)
		params = append(params, req.Text)
		paramCount++
	}
	if req.Link != "" {
		query += fmt.Sprintf(" link = $%d,", paramCount)
		params = append(params, req.Link)
		paramCount++
	}

	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE id = $%d", paramCount)
	params = append(params, songID)

	_, err := s.db.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update song: %v", err)
	}

	return nil
}
func (s *SongLibrary) GetSongText(songID string, verseNum int) (*models.VerseResponse, error) {
	query := `SELECT text FROM songs WHERE id = $1`
	var fullText string
	err := s.db.QueryRow(query, songID).Scan(&fullText)
	if err != nil {
		return nil, fmt.Errorf("failed to get song text: %v", err)
	}


	verses := strings.Split(fullText, "\n\n")


	if verseNum < 1 || verseNum > len(verses) {
		return nil, fmt.Errorf("verse number %d out of range, total verses: %d", verseNum, len(verses))
	}


	return &models.VerseResponse{
		Id:  songID,
		VersNum: verseNum,
		Text:    strings.TrimSpace(verses[verseNum-1]),
	}, nil
}


