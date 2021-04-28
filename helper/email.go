package helper

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type SMTPConnection struct {
	Host     string
	Sender   string
	Port     int
	Email    string
	Password string
}

func EmailConfig() SMTPConnection {
	port, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))

	dbConfig := SMTPConnection{
		Host:     os.Getenv("CONFIG_SMTP_HOST"),
		Port:     port,
		Sender:   os.Getenv("CONFIG_SENDER_NAME"),
		Email:    os.Getenv("CONFIG_AUTH_EMAIL"),
		Password: os.Getenv("CONFIG_AUTH_PASSWORD"),
	}

	return dbConfig
}

func SendEmail(recipient, subject string, body string) bool {
	c := EmailConfig()
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", c.Sender)
	mailer.SetHeader("To", recipient)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		c.Host,
		c.Port,
		c.Email,
		c.Password,
	)

	go dialer.DialAndSend(mailer)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return false
	// }

	log.Println("Mail sent!")
	return true
}
