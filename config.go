package main

import (
	"encoding/json"
	"os"
)

type Schedule struct {
	Time    string
	Zone    string
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
	NumStories int
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
	if checkErr("config.go: opening conf file", err) {
		return configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	checkErr("config.go: decoding conf struct", err)
	return configuration, err
}
