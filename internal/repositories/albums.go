package repositories

import (
	"errors"
	"example/web-service-gin/internal/models"
	"slices"
)

var (
	ErrNotFound = errors.New("not found")
)

// Our application is very simple
// our db interface maps one-to-one with service currently
type AlbumsRepository interface {
	Get() ([]models.Album, error)
	Create(models.Album) error
	GetById(id string) (models.Album, error)
}

type repo struct {
	albums []models.Album
}

func NewAlbumsRepository() *repo {
	return &repo{
		albums: slices.Clone(albums),
	}
}

func (r *repo) Get() ([]models.Album, error) {
	return r.albums, nil
}

func (r *repo) Create(newAlbum models.Album) error {
	r.albums = append(r.albums, newAlbum)
	return nil
}

func (r *repo) GetById(id string) (album models.Album, err error) {
	for _, a := range r.albums {
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
