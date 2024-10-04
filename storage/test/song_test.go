package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Online_Song_Libraries/config"
	"github.com/Online_Song_Libraries/models"
	"github.com/Online_Song_Libraries/storage/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var cnf = config.Load()

func TestAddSong(t *testing.T) {
	db, err := postgres.ConnectDB(cnf)
	assert.NoError(t, err)
	defer db.Close()

	songLibrary := postgres.NewSongLibrary(db)

	req := &models.AddSongRequest{
		GroupName:   "Muse",
		SongName:    "Supermassive Black Hole",
		ReleaseDate: "2006-07-16",
		Text:        "Ooh baby, don't you know I suffer?",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	res, err := songLibrary.AddSong(req)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.NotEmpty(t, res.SongID)
}

func TestGetAllSongs(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	songLibrary := postgres.NewSongLibrary(db)

	rows := sqlmock.NewRows([]string{"id", "group_name", "song_name", "release_date", "text", "link"}).
		AddRow(uuid.NewString(), "Muse", "Supermassive Black Hole", "2006-07-16", "Ooh baby, don't you know I suffer?", "https://www.youtube.com/watch?v=Xsp3_a-PMTw").
		AddRow(uuid.NewString(), "Queen", "Bohemian Rhapsody", "1975-10-31", "Is this the real life? Is this just fantasy?", "https://www.youtube.com/watch?v=fJ9rUzIMcZQ")

	mock.ExpectQuery("SELECT id, group_name, song_name, release_date, text, link FROM songs WHERE 1=1 LIMIT $1 OFFSET $2").
		WithArgs(10, 0).
		WillReturnRows(rows)

	songs, err := songLibrary.GetAll(&models.GetAllSongsRequest{Limit: 10, Offset: 0})
	assert.NoError(t, err)
	assert.Len(t, songs, 2)
	assert.Equal(t, songs[0].GroupName, "Muse")
	assert.Equal(t, songs[1].GroupName, "Queen")
}

func TestDeleteSong(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	songLibrary := postgres.NewSongLibrary(db)

	songID := uuid.NewString()

	mock.ExpectExec("DELETE FROM songs WHERE id = $1").
		WithArgs(songID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := songLibrary.DeleteSong(&models.DeleteSongRequest{Id: songID})
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "Song deleted successfully", res.Message)
}


func TestUpdateSong(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	songLibrary := postgres.NewSongLibrary(db)

	songID := uuid.NewString()

	mock.ExpectExec("UPDATE songs SET group_name = $1, song_name = $2, release_date = $3, text = $4, link = $5 WHERE id = $6").
		WithArgs("Muse", "Supermassive Black Hole", "2006-07-16", "Ooh baby, don't you know I suffer?", "https://www.youtube.com/watch?v=Xsp3_a-PMTw", songID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := songLibrary.UpdateSong(&models.UpdateSongRequest{
		Id:          songID,
		GroupName:   "Muse",
		SongName:    "Supermassive Black Hole",
		ReleaseDate: "2006-07-16",
		Text:        "Ooh baby, don't you know I suffer?",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	})
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "Song updated successfully", res.Message)
}

func TestGetSongText(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	songLibrary := postgres.NewSongLibrary(db)

	songID := uuid.NewString()
	fullText := "Verse 1\n\nVerse 2\n\nVerse 3"

	mock.ExpectQuery("SELECT text FROM songs WHERE id = $1").
		WithArgs(songID).
		WillReturnRows(sqlmock.NewRows([]string{"text"}).AddRow(fullText))

	res, err := songLibrary.GetSongText(&models.GetSongTextRequest{Id: songID, VerseNum: 2})
	assert.NoError(t, err)
	assert.Equal(t, songID, res.Id)
	assert.Equal(t, 2, res.VerseNum)
	assert.Equal(t, "Verse 2", res.Text)
}