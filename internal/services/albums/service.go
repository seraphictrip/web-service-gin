package albums

import (
	"errors"
	"example/web-service-gin/internal/models"
)

var (
	ErrNotFound = errors.New("not found")
)

type AlbumsService interface {
	Get() ([]models.Album, error)
	Create(models.Album) error
	GetById(id string) (models.Album, error)
}

type svc struct {
	// TODO: hook up a repository
}

func NewAlbumsService() *svc {
	return &svc{}
}
func (s *svc) Get() ([]models.Album, error) {
	return albums, nil
}

func (s *svc) Create(newAlbum models.Album) error {
	albums = append(albums, newAlbum)
	return nil
}

func (s *svc) GetById(id string) (album models.Album, err error) {
	for _, a := range albums {
		if a.ID == id {

			return a, nil
		}
	}
	return album, ErrNotFound
}

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
