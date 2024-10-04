package handler

import (
	"net/http"
	"strconv"

	"github.com/Online_Song_Libraries/models"
	"github.com/gin-gonic/gin"

)

// @Summary Add new song
// @Description Add new song to the library
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song body models.AddSongRequest true "Song data"
// @Success 200 {object} models.AddSongResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs [post]
func (h *Handler) AddSong(c *gin.Context) {
	var request models.AddSongRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.song.AddSong(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, res)

}

// @Summary Get all songs
// @Description Get all songs from the library
// @Tags songs
// @Accept  json
// @Produce  json
// @Param group_name query string false "Group name"
// @Param song_name query string false "Song name"
// @Param text query string false "Text"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} models.Song
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs/all [get]
func (h *Handler) GetAll(c *gin.Context) {
	var request models.GetAllSongsRequest

	request.GroupName = c.Query("group_name")
	request.SongName = c.Query("song_name")
	request.Text = c.Query("text")

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	if limitStr == "" {
		request.Limit = 10
	} else {
		limit, err := strconv.Atoi(limitStr)
		if err == nil {
			request.Limit = limit
		}
	}

	if offsetStr == "" {
		request.Offset = 0 
	} else {
		offset, err := strconv.Atoi(offsetStr)
		if err == nil {
			request.Offset = offset
		}
	}

	songs, err := h.song.GetAll(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// @Summary Delete song
// @Description Delete song by ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id query string true "Song ID"
// @Success 200 {object} models.DeleteSongResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs/{id} [delete]
func (h *Handler) DeleteSong(c *gin.Context) {
	songID := c.Query("id")
	res, err := h.song.DeleteSong(&models.DeleteSongRequest{Id: songID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}


// write swagger
// @Summary Update song
// @Description Update song by ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id query string true "Song ID"
// @Param group_name query string false "Group Name"
// @Param song_name query string false "Song Name"
// @Param release_date query string false "Release Date"
// @Param text query string false "Text"
// @Param link query string false "Link"
// @Success 200 {object} models.UpdateSongResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs/{id} [put]
func (h *Handler) UpdateSong(c *gin.Context) {
	request := models.UpdateSongRequest{}
	request.Id = c.Query("id")
	request.GroupName = c.Query("group_name")
	request.SongName = c.Query("song_name")
	request.ReleaseDate = c.Query("release_date")
	request.Text = c.Query("text")
	request.Link = c.Query("link")

	

	res, err := h.song.UpdateSong(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}



// @Summary Get song text by verse number
// @Description Get song text by verse number
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id query string true "Song ID"
// @Param verse_num query int true "Verse Number"
// @Success 200 {object} models.VerseResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs/{id}/verse/{verse_num} [get]
func (h *Handler) GetSongText(c *gin.Context) {
	request := models.GetSongTextRequest{}
	request.Id = c.Query("id")
	verseNumStr := c.Query("verse_num")
	if verseNumStr != ""{
		verseNum, err := strconv.Atoi(verseNumStr)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verse number"})
			return
		}
		request.VerseNum = verseNum
	} 


	verseResponse, err := h.song.GetSongText(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, verseResponse)
}
