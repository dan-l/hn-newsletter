# hn-newsletter-bot

* Automate retrieving of HN top stories daily
* Configurable to send newsletter at certain intervals and specify the number of stories to fetch in each newsletter

## Development

[Install Go](https://golang.org/doc/install)

[Install Govendor](https://github.com/kardianos/govendor)

Create configuration file  
`touch conf.json`

Format of conf.json
```
{
	"Mailgun": {
		"Sender": <sender email>,
		"Recipient": <recipient email>,
		"Schedules": [{
			"Time": "02 Jan 06 12:00 -0700",
		        "Zone": "America/Los_Angeles"
            	}],
		"Api": <mailgun api key>,
		"Domain": <email domain>,
		"Subject": <email subject>,
		"Body": <email body>
	},
	"Hn": {
		"Num_Stories": <number of stories to send in newsletter>
	}
}
```

Fetch package dependencies  
`govendor fetch +m`

Build binary  
`govendor build github.com/dan-l/hn-newsletter`

Run binary  
`./hn-newsletter`

## Third-party dependencies
1. [Mailgun](https://documentation.mailgun.com/quickstart.html), REST API used to send the email
2. [Firebase Hacker News API](https://hacker-news.firebaseio.com), REST API used to query stories
3. [GoCron](https://github.com/jasonlvhit/gocron), Cron job in Go

## Deployment with Heroku
Create app  
`heroku create`

Deploy and Build  
```
git add .
git commit -m <commit message> 
git push heroku master
```

Start worker  
`heroku ps:scale worker=1`

Check log  
`heroku logs --tail`
