package service

import (
	"context"

	"github.com/Dimix-international/readwise-go/db"
	"github.com/Dimix-international/readwise-go/internal/models"
)

type UsersService struct {
	bookStorage db.BookStorage
}

func NewUsersService(bookStorage db.BookStorage) *UsersService {
	return &UsersService{bookStorage: bookStorage}
}

func (s *UsersService) Users(ctx context.Context) ([]*models.User, error) {
	users, err := s.bookStorage.Users(ctx)
	if err != nil {
		return nil, models.InternalServerError
	}

	return users, nil
}
