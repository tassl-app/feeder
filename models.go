package feeder

import (
	"time"
)

const feedTimeLayout = "02 Jan 2006 15:04:05 -0700"

type Entry struct {
	MediaGroup     string   `json:"mediaGroup"`
	Title          string   `json:"title"`
	Link           string   `json:"link"`
	Content        string   `json:"content"`
	ContentSnippet string   `json:"contentSnippet"`
	PublishedDate  string   `json:"publishedDate"`
	Categories     []string `json:"categories"`
}

func (e *Entry) FormattedPublishedDate() (time.Time, error) {
	return time.Parse(feedTimeLayout, e.PublishedDate)
}

type Feed struct {
	FeedUrl     string  `json:"feedUrl"`
	Title       string  `json:"title"`
	Link        string  `json:"link"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	Entries     []Entry `json:"entries"`
}

type ResponseData struct {
	Feed Feed `json:"feed"`
}

type FeedResponse struct {
	ResponseData    ResponseData `json:"responseData"`
	ResponseDetails string       `json:"responseDetails"`
	ResponseStatus  int64        `json:"responseStatus"`
}
