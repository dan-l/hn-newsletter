package main

import (
	"log"
	"sync"
)

var waitGroup sync.WaitGroup // Sync variables.

func main() {
	configuration, err := readConf()
	if err != nil {
		return
	}
	formatSchedule(configuration.Mailgun.Schedules)
	log.Println(configuration.Mailgun)
	log.Println(configuration.Hn)

	waitGroup.Add(1)
	ScheduleHnNewsletterJob(configuration)
	waitGroup.Wait()
	log.Println("Done")
}

func checkErr(str string, err error) bool {
	if err != nil {
		log.Fatal(str, " ", err)
		return true
	}
	return false
}

func sendErr(conf Configuration, err error) {
	subject := conf.Mailgun.Subject
	var body string
	body = err.Error()
	sender := conf.Mailgun.Sender
	recipient := conf.Mailgun.Recipient
	api := conf.Mailgun.Api
	domain := conf.Mailgun.Domain

	_, err = SendSimpleMessage("", subject, body, sender, recipient, domain, api)
	if checkErr("app.go line 47: sendErr", err) {
		return
	}
}
