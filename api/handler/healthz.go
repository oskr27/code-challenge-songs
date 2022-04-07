package handler

import (
	"net/http"

	"github/CreatorNinja/code-challenge-songs/domain/healthz"

	"github.com/gin-gonic/gin"
)

func HealthzRouter(e *gin.Engine, s healthz.UseCase) {
	e.GET("/healthz", GetHealthz(s))
	
	// Here add all the methods for Create, Read, and Delete
	// Create methods
	 e.POST("/healthz/create-song", CreateSong(s))
	// e.POST("/healthz/create-artist", CreateArtist(s))
	// e.POST("/healthz/create-album", CreateAlbum(s))
	// e.POST("/healthz/create-playlist", CreatePlaylist(s))

	// Read Methods
	// e.GET("/healthz/list-song", GetListSong(s))
	// e.GET("/healthz/list-artist", GetListArtist(s))
	// e.GET("/healthz/list-album", GetListAlbum(s))
	// e.GET("/healthz/list-playlist", GetListPlaylist(s))

	// Delete Methods
	// e.DELETE("/healthz/delete-song", DeleteSong(s))
	// e.DELETE("/healthz/delete-artist", DeleteArtist(s))
	// e.DELETE("/healthz/delete-album", DeleteAlbum(s))
	// e.DELETE("/healthz/delete-playlist", DeletePlaylist(s))
}

func GetHealthz(s healthz.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		status, err := s.GetHealthz()
		
		if err != nil {
			c.JSON(http.StatusCreated, gin.H{
				"msg":    "Error on the request",
				"status": http.StatusBadRequest,
				"data":   nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":    status,
			"status": http.StatusOK,
			"data":   "OK",
		})
		return
	}
}

func CreateSong(s healthz.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		status, err := s.GetHealthz()
		
		if err != nil {
			c.JSON(http.StatusCreated, gin.H{
				"msg":    "Error on the request",
				"status": http.StatusBadRequest,
				"data":   nil,
			})
			return
		}
	
		//Otherwise, create song into the database
		name := c.Query("name")
		
		// Add code for inserting into database here
		status, err = s.CreateSong(name)
		c.JSON(http.StatusOK, gin.H{
			"msg":    status,
			"status": http.StatusOK,
			"data":   "OK",
		})
		return
	}
}
