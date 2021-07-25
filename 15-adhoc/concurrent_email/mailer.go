package main

import (
	"bufio"
	"fmt"
	"github.com/jordan-wright/email"
	clog "github.com/withmandala/go-log"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
	"time"
)

const (
	smtpHostAndPort  = "smtp.gmail.com:587"
	GMAIL_SMTP       = "smtp.gmail.com"
	REGISTRATION_MSG = "Welcome to our website !"
)

var email_ids = []string{"khabsdbhkzchkbdahd_01@yopmail.com", "khabsdbhkzchkbdahd_02@yopmail.com", "khabsdbhkzchkbdahd_03@yopmail.com", "khabsdbhkzchkbdahd_04@yopmail.com", "khabsdbhkzchkbdahd_05@yopmail.com", "khabsdbhkzchkbdahd_06@yopmail.com", "khabsdbhkzchkbdahd_07@yopmail.com", "khabsdbhkzchkbdahd_08@yopmail.com", "khabsdbhkzchkbdahd_09@yopmail.com", "khabsdbhkzchkbdahd_10@yopmail.com", "khabsdbhkzchkbdahd_11@yopmail.com", "khabsdbhkzchkbdahd_12@yopmail.com", "khabsdbhkzchkbdahd_13@yopmail.com", "khabsdbhkzchkbdahd_14@yopmail.com", "khabsdbhkzchkbdahd_15@yopmail.com", "khabsdbhkzchkbdahd_16@yopmail.com", "khabsdbhkzchkbdahd_17@yopmail.com", "khabsdbhkzchkbdahd_18@yopmail.com", "khabsdbhkzchkbdahd_19@yopmail.com", "khabsdbhkzchkbdahd_20@yopmail.com", "khabsdbhkzchkbdahd_21@yopmail.com", "khabsdbhkzchkbdahd_22@yopmail.com", "khabsdbhkzchkbdahd_23@yopmail.com", "khabsdbhkzchkbdahd_24@yopmail.com", "khabsdbhkzchkbdahd_25@yopmail.com", "khabsdbhkzchkbdahd_26@yopmail.com", "khabsdbhkzchkbdahd_27@yopmail.com", "khabsdbhkzchkbdahd_28@yopmail.com", "khabsdbhkzchkbdahd_29@yopmail.com", "khabsdbhkzchkbdahd_30@yopmail.com", "khabsdbhkzchkbdahd_31@yopmail.com", "khabsdbhkzchkbdahd_32@yopmail.com", "khabsdbhkzchkbdahd_33@yopmail.com", "khabsdbhkzchkbdahd_34@yopmail.com", "khabsdbhkzchkbdahd_35@yopmail.com", "khabsdbhkzchkbdahd_36@yopmail.com", "khabsdbhkzchkbdahd_37@yopmail.com", "khabsdbhkzchkbdahd_38@yopmail.com", "khabsdbhkzchkbdahd_39@yopmail.com", "khabsdbhkzchkbdahd_40@yopmail.com"}
var successChan chan string = make(chan string)
var errorChan chan string = make(chan string)
var logger *clog.Logger
var name, html_email, senderEmail, senderPassword string

func init() {
	// initialising logger
	logger = clog.New(os.Stderr).WithColor()

	// reading the generic html template
	bs, _ := ioutil.ReadFile("index.html")
	html_email = string(bs)
}

func getNameFromEmail(email string) (name string) {
	return strings.Split(email, "@")[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter sender's email address: ")
	scanner.Scan()
	senderEmail = scanner.Text()

	fmt.Print("Enter password to send email: ")
	scanner.Scan()
	senderPassword = scanner.Text()

	fmt.Println("start\t", time.Now().Format("01-02-2006 15:04:05"))
	for _, email := range email_ids {
		name = getNameFromEmail(email)
		/*
			If we dont add sleep here, we get following error:
			Temporary System Problem. Try again later (10). m16sm18984115pgj.26 - gsmtp
			uncoment followinf line to fix the issue:
		*/
		// time.Sleep(1500 * time.Millisecond)
		go sendEmail(name, email)
	}
	for i := 0; i < len(email_ids); i++ {
		select {
		case toEmail, ok := <-errorChan:
			if ok {
				logger.Error("Could not mail to", toEmail)
			}
		case toEmail, ok := <-successChan:
			if ok {
				logger.Info("Mailed succesfuly:", toEmail)
			}
		}
	}
	fmt.Println("end\t", time.Now().Format("01-02-2006 15:04:05"))
	fmt.Println("About to exit")
}

func sendEmail(name, toEmail string) {
	// setting up the body
	nameReplaced := strings.Replace(html_email, "${username}", name, -1)
	emailReplaced := strings.Replace(nameReplaced, "${email_id}", toEmail, -1)

	// setting up the smtp config
	e := email.NewEmail()
	e.From = senderEmail
	e.Subject = REGISTRATION_MSG
	e.AttachFile("test.txt")
	e.To = []string{toEmail}
	e.HTML = []byte(emailReplaced)
	err := e.Send(smtpHostAndPort, smtp.PlainAuth("", senderEmail, senderPassword, GMAIL_SMTP))

	if err != nil {
		fmt.Println(" *** Reason ***\n", err)
		errorChan <- toEmail
	} else {
		successChan <- toEmail
	}
}
