package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

type ConfigMail struct {
	host     string
	port     int
	user     string
	password string
}

func main() {
	mail := ConfigMail{
		host:     "example.smt.mail.com",
		port:     587,
		user:     "johndo@mail.com",
		password: "123456",
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", "johndo@mail.com")
	msg.SetHeader("To", "otherguy@mail.com")
	msg.SetHeader("Subject", "test message")
	msg.SetBody("text/html", "<h1> Body test </h1>")

	dialer := gomail.NewDialer(mail.host, mail.port, mail.user, mail.password)

	err := dialer.DialAndSend(msg)
	if err != nil {
		log.Printf("Error sending the email: %v\n", err)
	} else {
		log.Println("Email successfully sent")
	}

}
