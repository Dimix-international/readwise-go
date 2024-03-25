package service

import (
	"context"
	"mime/multipart"

	"github.com/Dimix-international/readwise-go/internal/models"
)

type ServiceFile interface {
	ParseKindleFile(ctx context.Context, file *multipart.File, userID string) error
	RandomHighlights(ctx context.Context, limit, userID int) ([]*models.Highlight, error)
	BookByISBN(ctx context.Context, ID string) (models.Book, error)
}

type ServiceUsers interface {
	Users(ctx context.Context) ([]*models.User, error)
}

type ServiceCloud interface {
	SendInsightsEmails(ctx context.Context) error
}

type ServiceMailer interface {
	SendInsights(ins []*models.DailyInsight, u *models.User) error
}
