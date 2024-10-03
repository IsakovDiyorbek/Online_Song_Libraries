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




// @Summary Get songs
// @Description Get songs from the library
// @Tags songs
// @Accept  json
// @Produce  json
// @Param group_name query string false "Group name"
// @Param song_name query string false "Song name"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {array} models.Song
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs [get]
func (h *Handler) GetSongs(c *gin.Context) {
	var filter models.SongFilter

	filter.GroupName = c.Query("group_name")
	filter.SongName = c.Query("song_name")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	if pageStr == "" {
		filter.Page = 1 
	} else {
		page, err := strconv.Atoi(pageStr)
		if err == nil {
			filter.Page = page
		}
	}

	if pageSizeStr == "" {
		filter.PageSize = 10 
	} else {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err == nil {
			filter.PageSize = pageSize
		}
	}

	songs, err := h.song.GetSongs(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}



// write swagger
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
// @Param id path int true "Song ID"
// @Success 200 {object} models.AddSongResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs/{id} [delete]
func (h *Handler) DeleteSong(c *gin.Context) {
	songID := c.Query("id")


	err := h.song.DeleteSong(songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
}


// write swagger
// @Summary Update song
// @Description Update song by ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path string true "Song ID"
// @Param song body models.UpdateSongRequest true "Song data"
// @Success 200 {object} models.AddSongResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs/{id} [put]
func (h *Handler) UpdateSong(c *gin.Context) {
	songID:= c.Query("id")

	var request models.UpdateSongRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.song.UpdateSong(songID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully"})
}



// @Summary Get song text by verse number
// @Description Get song text by verse number
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path string true "Song ID"
// @Param verseNum path int true "Verse number"
// @Success 200 {object} models.VerseResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /songs/{id}/verse/{verseNum} [get]
func (h *Handler) GetSongText(c *gin.Context) {

	songIDQuery := c.Query("id")
	verseNumParam := c.Query("verseNum")

	verseNum, err := strconv.Atoi(verseNumParam)
	if err != nil || verseNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verse number"})
		return
	}


	verseResponse, err := h.song.GetSongText(songIDQuery, verseNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, verseResponse)
}
