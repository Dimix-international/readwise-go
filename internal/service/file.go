package service

import "mime/multipart"

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

func (s *FileService) ParseKindleFile(file *multipart.File, userID string) error {
	return nil
}
