package db

import (
	"database/sql"

	"github.com/Dimix-international/readwise-go/internal/models"
)

type BookMySQLStorage struct {
	db *sql.DB
}

func NewBookStorage(db *sql.DB) *BookMySQLStorage {
	return &BookMySQLStorage{db: db}
}

func (b *BookMySQLStorage) CreateBook(models.Book) error {
	return nil
}

func (b *BookMySQLStorage) CreateHighlights([]models.Highlight) error {
	return nil
}

func (b *BookMySQLStorage) BookByISBN(string) (models.Book, error) {
	return models.Book{}, nil
}
