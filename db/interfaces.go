package db

import (
	"context"

	"github.com/Dimix-international/readwise-go/internal/models"
)

type BookStorage interface {
	CreateBook(ctx context.Context, book models.Book) error
	CreateHighlights(ctx context.Context, highlights []models.Highlight) error
	BookByISBN(ctx context.Context, ID string) (models.Book, error)
	RandomHighlights(ctx context.Context, limit, userID int) ([]*models.Highlight, error)
	Users(ctx context.Context) ([]*models.User, error)
}
