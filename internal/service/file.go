package service

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/Dimix-international/readwise-go/db"
	"github.com/Dimix-international/readwise-go/internal/models"
)

type FileService struct {
	bookStorage db.BookStorage
}

func NewFileService(bookStorage db.BookStorage) *FileService {
	return &FileService{bookStorage: bookStorage}
}

func (s *FileService) ParseKindleFile(ctx context.Context, file *multipart.File, userID string) error {
	raw, err := s.parseKindleExtractFile(file)
	if err != nil {
		return models.InternalServerError
	}

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		return models.InternalServerError
	}

	if err := s.createDataFromRawBook(ctx, raw, userIDint); err != nil {
		return models.InternalServerError
	}

	return nil
}

func (s *FileService) RandomHighlights(ctx context.Context, limit, userID int) ([]*models.Highlight, error) {
	hs, err := s.bookStorage.RandomHighlights(ctx, limit, userID)
	if err != nil {
		return nil, err
	}
	return hs, nil
}

func (s *FileService) BookByISBN(ctx context.Context, ID string) (models.Book, error) {
	book, err := s.bookStorage.BookByISBN(ctx, ID)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s *FileService) parseKindleExtractFile(file *multipart.File) (*models.RawExtractBook, error) {
	decoder := json.NewDecoder(*file)

	raw := models.RawExtractBook{}
	if err := decoder.Decode(&raw); err != nil {
		return nil, err
	}

	return &raw, nil
}

func (s *FileService) createDataFromRawBook(ctx context.Context, raw *models.RawExtractBook, userID int) error {
	if _, err := s.bookStorage.BookByISBN(ctx, raw.ASIN); err != nil {
		s.bookStorage.CreateBook(ctx, models.Book{
			ISBN:      raw.ASIN,
			Title:     raw.Title,
			Authors:   raw.Authors,
			CreatedAt: time.Now().UTC(),
		})
	}

	hs := make([]models.Highlight, len(raw.Highlights))
	for i := range raw.Highlights {
		hs[i] = models.Highlight{
			Text:     raw.Highlights[i].Text,
			Location: raw.Highlights[i].Location.URL,
			Note:     raw.Highlights[i].Note,
			UserID:   userID,
			BookID:   raw.ASIN,
		}
	}

	if err := s.bookStorage.CreateHighlights(ctx, hs); err != nil {
		return err
	}

	return nil
}
