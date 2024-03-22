package service

import (
	"encoding/json"
	"mime/multipart"

	"github.com/Dimix-international/readwise-go/db"
	"github.com/Dimix-international/readwise-go/internal/models"
)

type FileService struct {
	bookStorage db.BookStorage
}

func NewFileService(bookStorage db.BookStorage) *FileService {
	return &FileService{bookStorage: bookStorage}
}

func (s *FileService) ParseKindleFile(file *multipart.File, userID string) error {
	_, err := s.parseKindleExtractFile(file)
	if err != nil {
		return models.InternalServerError
	}

	return nil
}

func (s *FileService) parseKindleExtractFile(file *multipart.File) (*models.RawExtractBook, error) {
	decoder := json.NewDecoder(*file)

	raw := models.RawExtractBook{}
	if err := decoder.Decode(&raw); err != nil {
		return nil, err
	}

	return &raw, nil
}
