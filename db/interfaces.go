package db

import "github.com/Dimix-international/readwise-go/internal/models"

type BookStorage interface {
	CreateBook(models.Book) error
	CreateHighlights([]models.Highlight) error
	BookByISBN(string) (models.Book, error)
}
