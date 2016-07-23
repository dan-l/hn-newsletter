# hn-newsletter-bot

* Automate retrieving of HN top stories daily
* Configurable to send newsletter at certain intervals and specify the number of stories to fetch in each newsletter

## Development

[Install Go](https://golang.org/doc/install)

[Install Go package manager](https://github.com/gpmgo/gopm/blob/master/README.md#installation)

Change directory to "server"  
`cd server`

Create configuration file  
`touch conf.json`

Format of conf.json
```
{
	"Mailgun": {
		"Sender": <sender email>,
		"Recipient": <recipient email>,
		"Schedules": [
	        {
	            "Time": "Mon Jan  1 09:00:00 PDT 2015",
	            "Zone": "America/Los_Angeles"
            },
            {
			    "Time": <some other time following the format above>,
                "Zone": <some other timezone follwing the format above>
            }
		],
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
`gopm get`

Build binary  
`gopm install`

Run binary  
`.vendor/bin/server`

## Third-party dependencies
1. [Mailgun](https://documentation.mailgun.com/quickstart.html), REST API used to send the email
2. [Firebase Hacker News API](https://hacker-news.firebaseio.com), REST API used to query stories
3. [GoCron](https://github.com/jasonlvhit/gocron), Cron job in Go


