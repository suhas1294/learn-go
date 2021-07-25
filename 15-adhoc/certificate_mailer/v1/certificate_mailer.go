package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	//"net/smtp"
	"encoding/csv"
	"io"
	"os"
)

var err error
var failedEmailList []string
var successIdList []string
var from_email string = ""

var common_mail_body string = `<!DOCTYPE html>
<html>
<head>
	<title></title>
</head>
<body>
<pre>
<b><span style="color: #0000ff; font-size: 20px;font-family: 'Arial, Helvetica, sans-serif;'" >Dear participant,</span></b>

<span style="color:#0000ff; font-size: 18px;font-family: 'Arial, Helvetica, sans-serif;'">Congratulations on the successful completion of the training programme.</span>
<span style="font-size: 15px;font-family: 'Arial, Helvetica, sans-serif;'">
We hope this training program may enrich the knowledge of XYZ field.
All the best for your future. Thank you for participation.
</span>

<b><span style="font-size: 18px;font-family: 'Arial, Helvetica, sans-serif;'">CUK-SIUD Team</span><b>
<b><span style="color: red;font-size: 20px;font-family: 'Arial, Helvetica, sans-serif;'">Please Do not reply to this mail.</span></b>
</pre>
</body>
</html>`

func main() {
	log.Println("starting the ops...")

	csv_file, err := os.Open("data.csv")
	defer csv_file.Close()
	checkError(err)
	reader := csv.NewReader(csv_file)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		checkError(err)
		sendMail(line[0], line[1]) // line[0] = receipients email address, line[1] = receipient's certificate name located in same folder
	}

	log.Println("\n\nSuccessfully sent to :", successIdList)
	log.Println("\n\nEnding the ops... Failed to send  email to following emails:", failedEmailList)
}

func sendMail(toEmail string, attachmentName string) {
	e := email.NewEmail()
	e.From = from_email
	e.To = []string{toEmail}
	e.Subject = "Successful completetion of online QGIS training"
	e.HTML = []byte(common_mail_body)
	e.AttachFile(attachmentName)
	//err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "example@gmail.com", "password", "smtp.gmail.com"))
	_, err = fmt.Println(toEmail, attachmentName)

	commonMsg := toEmail + "," + attachmentName + "\n"

	if err != nil {
		failedEmailList = append(failedEmailList, toEmail)
		f, err := os.OpenFile("failure.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		checkError(err)
		defer f.Close()
		_, err = f.WriteString(commonMsg)
		checkError(err)

	} else {
		successIdList = append(successIdList, attachmentName)
		f, err := os.OpenFile("success.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		checkError(err)
		defer f.Close()
		_, err = f.WriteString(commonMsg)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
