package service

import (
	"context"

	"github.com/Dimix-international/readwise-go/internal/models"
)

type CloudService struct {
	fileService   ServiceFile
	usersService  ServiceUsers
	mailerService ServiceMailer
}

func NewCloudService(fileService ServiceFile, usersService ServiceUsers, mailerService ServiceMailer) *CloudService {
	return &CloudService{fileService: fileService, usersService: usersService, mailerService: mailerService}
}

func (s *CloudService) SendInsightsEmails(ctx context.Context) error {
	users, err := s.usersService.Users(ctx)
	if err != nil {
		return models.InternalServerError
	}

	for i := range users {
		hs, err := s.fileService.RandomHighlights(ctx, 3, users[i].ID)
		if err != nil {
			return err
		}

		if len(hs) == 0 {
			continue
		}

		insights, err := s.buildInsights(ctx, hs)
		if err != nil {
			return err
		}

		err = s.mailerService.SendInsights(insights, users[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CloudService) buildInsights(ctx context.Context, hs []*models.Highlight) ([]*models.DailyInsight, error) {
	var insights []*models.DailyInsight

	for i := range hs {
		book, err := s.fileService.BookByISBN(ctx, hs[i].BookID)
		if err != nil {
			return nil, err
		}

		insights = append(insights, &models.DailyInsight{
			Text:        hs[i].Text,
			Note:        hs[i].Note,
			BookAuthors: book.Authors,
			BookTitle:   book.Title,
		})
	}

	return insights, nil
}
