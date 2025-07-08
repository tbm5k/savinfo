package mailer

import (
	"fmt"
	"net/smtp"
)

type Mailer struct {
	host string
	port int
	username string
	password string
	from string
}

type Email struct {
	To string
	Subject string
	Body string
}

func NewMailer(username, password, from string) *Mailer {
	return &Mailer{
		host: "smtp.gmail.com",
		port: 587,
		username: username,
		password: password,
		from: from,
	}
}

func (m *Mailer) Send(e *Email) error {
	addr := fmt.Sprintf("%v:%v", m.host, m.port)
	message := []byte(
		"From: " + m.from + "\r\n" +
		"To: " + e.To + "\r\n" +
		"Subject: " + e.Subject + "\r\n" +
		"\r\n" +
		e.Body)

	auth := smtp.PlainAuth("", m.username, m.password, m.host)

	err := smtp.SendMail(addr, auth, m.from, []string{e.To}, message)
	if err != nil {
		return err
	}

	return nil
}

