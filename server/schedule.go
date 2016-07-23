package main

import (
	"github.com/jasonlvhit/gocron"
	"log"
	"time"
)

func SendHnNewsletter(conf Configuration) {
	log.Println("SendHnNewsletter", time.Now().Format(time.UnixDate))
	html, err := ConstructHnNewsletter(conf.Hn.Num_Stories)
	if err != nil {
		sendErr(conf, err)
	}
	subject := conf.Mailgun.Subject
	body := conf.Mailgun.Body
	sender := conf.Mailgun.Sender
	recipient := conf.Mailgun.Recipient
	api := conf.Mailgun.Api
	domain := conf.Mailgun.Domain

	_, err = SendSimpleMessage(html, subject, body, sender, recipient, domain, api)
	if checkErr("schedule.go line 34: SendHnNewsletter", err) {
		return
	}
}

func ScheduleHnNewsletterJob(conf Configuration) {
	defer waitGroup.Done()
	schedules := conf.Mailgun.Schedules
	for _, schedule := range schedules {
		gocron.Every(1).Day().At(schedule.JobTime).Do(SendHnNewsletter, conf)
	}
	<-gocron.Start()
}
