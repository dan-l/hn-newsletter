package main

import (
	"log"
	"sync"
	"github.com/jasonlvhit/gocron"
)

var waitGroup sync.WaitGroup // Sync variables.

func main() {
	configuration, err := readConf()
	if err != nil {
		return
	}
	log.Println(configuration.Mailgun)
	log.Println(configuration.Hn)

	waitGroup.Add(1)
	ScheduleHnNewsletterJob(configuration)
	waitGroup.Wait()
	log.Println("Done")
}

func ScheduleHnNewsletterJob(conf Configuration) {
	defer waitGroup.Done()
	schedules := conf.Mailgun.Schedules
	for _, schedule := range schedules {
		jobTime := formatScheduleTime(schedule.Zone, schedule.Time)
		gocron.Every(1).Thursday().At(jobTime).Do(SendHnNewsletter, conf)
	}
	_, time := gocron.NextRun()
	log.Println(time)
	<-gocron.Start()
}
