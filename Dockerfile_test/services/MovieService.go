package services

import "Dockerfile_test/datamodels"

type MovieService interface {
	GetAll() []datamodels.Movie
	GetById(id int64) (datamodels.Movie, bool)
	DeleteById(id int64) bool
	UpdatePosterAndGenreById(id int64, poster string, genre string) (datamodels.Movie, error)
}

