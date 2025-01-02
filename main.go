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
	albums := router.Group("/albums")
	albumsHandlers.RegisterRoutes(albums)

	router.Run("localhost:8080")
}
