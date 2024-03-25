package service

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/Dimix-international/readwise-go/internal/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const FromName = "Notebase"

type SendGridMailer struct {
	FromEmail string
	Client    *sendgrid.Client
}

func NewSendGridMailer(apiKey, fromEmail string) *SendGridMailer {
	client := sendgrid.NewSendClient(apiKey)
	if client == nil {
		return nil
	}

	return &SendGridMailer{
		FromEmail: fromEmail,
		Client:    client,
	}
}

func (m *SendGridMailer) SendInsights(insights []*models.DailyInsight, u *models.User) error {
	if m.Client == nil {
		return fmt.Errorf("error API KEY")
	}

	if u.Email == "" {
		return fmt.Errorf("user has no email")
	}

	from := mail.NewEmail(FromName, m.FromEmail)
	userName := fmt.Sprintf("%v %v", u.FirstName, u.LastName)

	to := mail.NewEmail(userName, u.Email)

	html := BuildInsightsMailTemplate(u, insights)

	message := mail.NewSingleEmail(from, "Daily Insight(s)", to, "", html)
	if _, err := m.Client.Send(message); err != nil {
		return err
	}

	return nil
}

func BuildInsightsMailTemplate(u *models.User, ins []*models.DailyInsight) string {
	templ, err := template.ParseFiles("daily.templ")
	if err != nil {
		panic(err)
	}

	payload := struct {
		User     *models.User
		Insights []*models.DailyInsight
	}{
		User:     u,
		Insights: ins,
	}

	var out bytes.Buffer
	err = templ.Execute(&out, payload)
	if err != nil {
		panic(err)
	}

	return out.String()
}
