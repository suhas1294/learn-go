package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	// "net/http/httputil"
	"sync"
	"time"

	"github.com/jordan-wright/email"
)

type Center struct {
	CenterID     int       `json:"center_id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	StateName    string    `json:"state_name"`
	DistrictName string    `json:"district_name"`
	BlockName    string    `json:"block_name"`
	Pincode      int       `json:"pincode"`
	Lat          int       `json:"lat"`
	Long         int       `json:"long"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	FeeType      string    `json:"fee_type"`
	Sessions     []Session `json:"sessions"`
}

type Session struct {
	SessionID              string   `json:"session_id"`
	Date                   string   `json:"date"`
	AvailableCapacity      int      `json:"available_capacity"`
	MinAgeLimit            int      `json:"min_age_limit"`
	Vaccine                string   `json:"vaccine"`
	Slots                  []string `json:"slots"`
	AvailableCapacityDose1 int      `json:"available_capacity_dose1"`
	AvailableCapacityDose2 int      `json:"available_capacity_dose2"`
}

type Response struct {
	Centers []Center `json:"centers"`
}

const (
	// VACCINE_NAME = "COVISHIELD"
	VACCINE_NAME = "COVAXIN"
	DISTRICT_ID  = "266"
	AGE_LIMIT    = 18
	url          = `https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict`

	// mail communication
	senderEmail     = "example@gmail.com"
	senderPassword  = "example"
	smtpHostAndPort = "smtp.gmail.com:587"
	GMAIL_SMTP      = "smtp.gmail.com"
	ALERT_MSG       = "Covaxin alert !"
)

var TO_EMAIL = []string{"target1@example.com", "target2@example.com"}

var wait_group sync.WaitGroup

func main() {
	// part-1 date calculations
	var dates []string
	var responses []Response
	var email_body []string

	resultChannel := make(chan Response)

	today := time.Now().Format("02-01-2006")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("02-01-2006")
	day_after_tomo := time.Now().AddDate(0, 0, 2).Format("02-01-2006")

	dates = append(dates, today)
	dates = append(dates, tomorrow)
	dates = append(dates, day_after_tomo)

	wait_group.Add(len(dates))

	for _, date := range dates {
		go makeRequest(date, resultChannel)
	}

	go func() {
		for v := range resultChannel {
			responses = append(responses, v)
		}
	}()
	wait_group.Wait()
	fmt.Println("Response length:\t", len(responses))
	for _, value := range responses {
		for _, center := range value.Centers {
			for _, session := range center.Sessions {
				if session.Vaccine == VACCINE_NAME && session.MinAgeLimit == AGE_LIMIT && session.AvailableCapacityDose1 > 0 {
					slots := strings.Join(session.Slots[:], "\t\t")
					mailBody := fmt.Sprintf("Date:\t\t%v\nCenter Name:\t%v\nCenter addr:\t%v\nPincode:\t%v\nSlots:\t\t%v\ncapicity:\t%d\n\n\n", session.Date, center.Name, center.Address, center.Pincode, slots, session.AvailableCapacity)
					fmt.Println(mailBody)
					email_body = append(email_body, mailBody)
				}
			}
		}
	}

	if len(email_body) > 0 {
		sendEmail(email_body)
	}

	f, err := os.OpenFile("vaccine_cron.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	timeNow := time.Now()
	_, err = f.Write([]byte(timeNow.String() + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	fmt.Println("\n\t\t********** Execution Ended **********")

}

func makeRequest(date string, resChan chan Response) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic("Unable to construct request")
	}

	q := req.URL.Query()
	q.Add("district_id", DISTRICT_ID)
	q.Add("date", date)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("accept", "application/json")
	req.Header.Add("Accept-Language", "hi_IN")
	req.Header.Add("User-Agent", "PostmanRuntime/7.26.10")

	// making the request
	client := &http.Client{}
	resp, err := client.Do(req)

	// interceptor.FormatRequest(req)
	// requestDump, _ := httputil.DumpRequest(req, true)
	// fmt.Println(string(requestDump))

	if err != nil {
		fmt.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// converting response to struct
	var response Response
	if err = json.Unmarshal(respBody, &response); err != nil {
		fmt.Println(err)
	}
	resChan <- response
	wait_group.Done()
}

func sendEmail(mailBody []string) {
	body := strings.Join(mailBody, "\n\n")
	e := email.NewEmail()
	e.From = senderEmail
	e.Subject = ALERT_MSG
	e.To = TO_EMAIL
	e.Text = []byte(body)
	err := e.Send(smtpHostAndPort, smtp.PlainAuth("", senderEmail, senderPassword, GMAIL_SMTP))

	if err != nil {
		fmt.Println("Sending mail failed", err)
	} else {
		fmt.Println("Sent email successfully")
	}
}
