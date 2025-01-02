package handlers

import (
	"errors"
	"example/web-service-gin/internal/models"
	albums_svc "example/web-service-gin/internal/services/albums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc albums_svc.AlbumsService
}

func NewAlbumsHandler(svc albums_svc.AlbumsService) *Handler {
	return &Handler{
		svc: svc,
	}
}

// albums slice to seed record album data.
// var albums = []models.Album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// getAlbums responds with the list of all albums as JSON.
func (h *Handler) GetAlbums(c *gin.Context) {
	albums, err := h.svc.Get()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func (h *Handler) PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	if err := h.svc.Create(newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (h *Handler) GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	album, err := h.svc.GetById(id)
	if err != nil {
		if errors.Is(err, albums_svc.ErrNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
	return

}

func GetPath(c *gin.Context) {
	path := c.Param("path")

	c.IndentedJSON(http.StatusOK, gin.H{"path": path})
}
