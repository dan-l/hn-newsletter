package main

import (
	"github.com/mailgun/mailgun-go"
)

func SendSimpleMessage(html, subject, body, sender, recipient, domain, api string) (string, error) {
	mg := mailgun.NewMailgun(domain, api, "")
	m := mg.NewMessage(
		sender,
		subject,
		body,
		recipient,
	)
	if html != "" {
		m.SetHtml(html)
	}
	_, id, err := mg.Send(m)
	return id, err
}
