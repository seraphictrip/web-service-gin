package main

import (
	"example/web-service-gin/internal/handlers"
	"example/web-service-gin/internal/repositories"
	albums_svc "example/web-service-gin/internal/services/albums"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	albumRepo := repositories.NewAlbumsRepository()
	albumsService := albums_svc.NewAlbumsService(albumRepo)
	albumsHandlers := handlers.NewAlbumsHandler(albumsService)

	// Albums
	router.GET("/albums", albumsHandlers.GetAlbums)
	router.POST("/albums", albumsHandlers.PostAlbums)

	router.GET("/albums/:id", albumsHandlers.GetAlbumById)

	router.Run("localhost:8080")
}
