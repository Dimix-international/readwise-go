package db

import (
	"context"
	"database/sql"

	"github.com/Dimix-international/readwise-go/internal/models"
)

type BookMySQLStorage struct {
	db *sql.DB
}

func NewBookStorage(db *sql.DB) *BookMySQLStorage {
	return &BookMySQLStorage{db: db}
}

func (s *BookMySQLStorage) CreateBook(ctx context.Context, b models.Book) error {
	if _, err := s.db.ExecContext(ctx, `
		INSERT INTO books (isbn, title, authors)
		VALUES (?, ?, ?, ?)
	`, b.ISBN, b.Title, b.Authors, b.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (s *BookMySQLStorage) CreateHighlights(ctx context.Context, hs []models.Highlight) error {
	values := []interface{}{}

	query := "INSERT INTO highlights (text, location, note, userId, bookId) VALUES "

	for i := range hs {
		query += "(?, ?, ?, ?, ?),"
		values = append(values, hs[i].Text, hs[i].Location, hs[i].Note, hs[i].UserID, hs[i].BookID)
	}

	query = query[:len(query)-1]

	_, err := s.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookMySQLStorage) BookByISBN(ctx context.Context, isbn string) (models.Book, error) {
	book := models.Book{}

	if err := s.db.QueryRowContext(ctx, `SELECT * FROM books WHERE isbn = ?`, isbn).Scan(&book); err != nil {
		return models.Book{}, models.NotFoundError
	}

	return book, nil
}

func (s *BookMySQLStorage) RandomHighlights(ctx context.Context, limit, userID int) ([]*models.Highlight, error) {
	rows, err := s.db.Query("SELECT * FROM highlights WHERE userId = ? ORDER BY RAND() LIMIT ?", userID, limit)
	if err != nil {
		return nil, err
	}

	var highlights []*models.Highlight
	for rows.Next() {
		h := new(models.Highlight)

		if err := rows.Scan(h); err != nil {
			return nil, err
		}

		highlights = append(highlights, h)
	}

	return highlights, nil
}

func (s *BookMySQLStorage) Users(ctx context.Context) ([]*models.User, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0)
	for rows.Next() {
		u := new(models.User)

		if err := rows.Scan(u); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
