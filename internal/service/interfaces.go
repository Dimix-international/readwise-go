package service

import (
	"context"
	"mime/multipart"
)

type ServiceFile interface {
	ParseKindleFile(ctx context.Context, file *multipart.File, userID string) error
}
