package main

import (
	"encoding/json"
	"os"
)

type Schedule struct {
	Time    string
	Zone    string
	JobTime string
}

type MailgunConf struct {
	Sender    string
	Recipient string
	Schedules []Schedule
	Api       string
	Domain    string
	Subject   string
	Body      string
}

type HnConf struct {
	Num_Stories int
}

type Configuration struct {
	Mailgun MailgunConf
	Hn      HnConf
}

const (
	CONF_FILE string = "conf.json"
)

func readConf() (Configuration, error) {
	configuration := Configuration{}
	file, err := os.Open(CONF_FILE)
	if checkErr("app.go line 40: opening conf file", err) {
		return configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	checkErr("config.go line 46: decoding conf struct", err)
	return configuration, err
}
