package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type HnStory struct {
	Url         string
	Title       string
	Score       int
	Time        int64
	Type        string
	CommentUrl  string
	Descendants int
}

const (
	HN_TOP_STORIES string = "https://hacker-news.firebaseio.com/v0/topstories.json"
	HN_STORY       string = "https://hacker-news.firebaseio.com/v0/item/"
	HN_POST        string = "https://news.ycombinator.com/item?id="
)

func JsonGet(url string, val interface{}) error {
	resp, err := http.Get(url)
	if checkErr("hn.go line 21: HttpGet "+url, err) {
		return err
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&val)
	if checkErr("hn.go line 24: JsonGet "+url, err) {
		return err
	}
	return nil
}

func GetHnStories(num_stories int) ([]HnStory, error) {
	var topstories []int
	err := JsonGet(HN_TOP_STORIES, &topstories)
	topstories = topstories[:num_stories]
	stories := make([]HnStory, num_stories)
	for i, storyid := range topstories {
		idstr := strconv.Itoa(storyid)
		var story HnStory
		JsonGet(HN_STORY+idstr+".json", &story)
		story.CommentUrl = HN_POST + idstr
		stories[i] = story
	}
	return stories, err
}

func ConstructHnNewsletter(num_stories int) (string, error) {
	stories, err := GetHnStories(num_stories)

	htmlString := "<header>" +
		"<div style='font-family: Verdana; font-size: 13.33px; color: black; text-decoration: none; background-color: #ff6600'>" +
		"<a href='https://news.ycombinator.com/news' style='color: black'><b>Hacker Newsletter</b></a>" +
		"</div>" +
		"</header>"
	htmlString += "<ul style='background-color: #f6f6ef; list-style:none'>"
	for _, story := range stories {
		storyEl := "<li style='padding: 15px'>" +
			"<span><a href='" + story.Url + "' style='color: black; text-decoration: none'>" + story.Title + "</a></span>" +
			"<div><small style='color: #666d74'>" +
			"<span style='padding-right: 10px'>" + strconv.Itoa(story.Score) + " points </span>" +
			"<span style='padding-right: 10px'>" + story.Type + "</span>" +
			"<span style='padding-right: 10px'>" + timeAgo(story.Time) + " </span>" +
			"<a href='" + story.CommentUrl + "' style='color: #666d74; text-decoration: none'>" +
			"|   " + strconv.Itoa(story.Descendants) + " Comments" +
			"</a>" +
			"</small></div>" +
			"</li>"

		htmlString += storyEl
	}
	htmlString += "</ul>"
	return htmlString, err
}
