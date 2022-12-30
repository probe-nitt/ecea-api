package services

import (
	"bytes"
	"crypto/tls"

	"github.com/ecea-nitt/ecea-server/config"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/schemas"
	"github.com/ecea-nitt/ecea-server/utils"
	"gopkg.in/gomail.v2"
)

type mailService struct{}

type MailService interface {
	MailUser(user schemas.User, emailData models.EmailData) error
}

func NewMailService() MailService {
	return &mailService{}
}

func (us *mailService) MailUser(user schemas.User, emailData models.EmailData) error {

	// Sender data.
	from := "ecea@nitt.edu"
	smtpPass := config.MailPassword
	smtpUser := config.MailUser
	to := user.Email
	smtpHost := config.MailHost
	smtpPort := config.MailPort

	var body bytes.Buffer

	template, err := utils.ParseTemplateDir("templates")
	if err != nil {
		return err
	}

	err = template.ExecuteTemplate(&body, "verificationCode.html", emailData)

	if err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", emailData.Subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
