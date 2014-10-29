package feeder

import (
	"testing"
)

const testRssFeed = "http://www.dailyherald.com/rss/feed/?feed=sports_all"

func TestAPIFeedLoad(t *testing.T) {
	api := NewAPI()
	feedResponse, err := api.FeedLoad(testRssFeed)
	if err != nil {
		t.Errorf("Error loading feed\n%s\n", err.Error())
		return
	}
	if feedResponse.ResponseStatus != 200 {
		t.Errorf("Expected response status %d, found %d\n", 200, feedResponse.ResponseStatus)
		return
	}
}
