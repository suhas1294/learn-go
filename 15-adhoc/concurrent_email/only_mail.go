package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"net/smtp"
)

var html_email string

const (
	senderEmail    = "example@gmail.com"
	senderPassword = "password"
	smtpHostAndPort  = "smtp.gmail.com:587"
	GMAIL_SMTP       = "smtp.gmail.com"
	REGISTRATION_MSG = "Welcome to our website !"
)

func init() {
	bs, _ := ioutil.ReadFile("index.html")
	html_email = string(bs)
}

func main() {
	c := make(chan string)
	go sendEmail(c)
	fmt.Println(<-c)
	fmt.Println("About to exit")
}

func sendEmail(c chan string) {
	e := email.NewEmail()
	e.From = senderEmail
	e.Subject = REGISTRATION_MSG
	e.AttachFile("test.txt")
	e.To = []string{"receiver@gmail.com"}
	e.HTML = []byte(html_email)
	err := e.Send(smtpHostAndPort, smtp.PlainAuth("", senderEmail, senderPassword, GMAIL_SMTP))

	if err != nil {
		c <- "Failed!"
	} else {
		c <- "Success!"
	}
}
