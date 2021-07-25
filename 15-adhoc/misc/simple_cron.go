package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("cron_test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	timeNow := time.Now()

	_, err = f.Write([]byte(timeNow.String() + "\n"))
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}

// 0/5 * * * * ? * ./Users/<username>/workspace/utils/go_cron
// 0/5 * * * * ? * "cd /Users/<username>/workspace/utils/crons && ./go_cron"
