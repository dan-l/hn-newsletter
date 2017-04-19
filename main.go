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
