package handler

import (
	"net/http"

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
