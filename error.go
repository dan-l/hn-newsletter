package main

import (
	"log"
)

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