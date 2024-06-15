package dao

import (
	"context"
	"github/MovieWebsite/global/db"
	"github/MovieWebsite/internal/model"
	"gorm.io/gorm"
)

type MovieDao struct {
	movie model.Movie
	db    *gorm.DB
}

func NewMovieDao(ctx context.Context, movie model.Movie) *MovieDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &MovieDao{
		movie: movie,
		db:    db.NewDBClient(ctx),
	}
}
