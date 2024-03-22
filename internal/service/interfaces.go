package service

import "mime/multipart"

type ServiceFile interface {
	ParseKindleFile(file *multipart.File, userID string) error
}
