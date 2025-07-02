package services

import (
	"bytes"
	"text/template"

	"github.com/rayhan889/talkz-v2/config"
	"gopkg.in/gomail.v2"
)

type MailService struct {
	dialer *gomail.Dialer
}

func NewMailService(dialer *gomail.Dialer) *MailService {
	return &MailService{
		dialer: dialer,
	}
}

func (service *MailService) SendMail(
	to string,
	subject string,
	templateFile string,
	data interface{},
) error {
	var body bytes.Buffer
	t, err := template.New("mailer").Parse(templateFile)

	if err != nil {
		return err
	}

	err = t.Execute(&body, data)

	if err != nil {
		return err
	}

	html := body.String()

	mailer := gomail.NewMessage()

	mailer.SetHeader("From", config.Mail.SenderName)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", html)

	return service.dialer.DialAndSend(mailer)
}
