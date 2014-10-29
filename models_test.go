package feeder

import (
	"encoding/json"
	"testing"
)

const feedResponseJSON = `
	{ "responseData" : { "feed" : { "author" : "",
	          "description" : "",
	          "entries" : [ { "author" : "",
	                "categories" : [  ],
	                "content" : "Bulls guard/forward Jimmy Butler is out for tonight's opening game against Carmelo Anthony and the New York Knicks. When he'll return from a sprained left thumb remains unclear.",
	                "contentSnippet" : "Bulls guard/forward Jimmy Butler is out for tonight's opening game against Carmelo Anthony and the New York Knicks. When he'll ...",
	                "link" : "http://www.dailyherald.com/article/20141029/sports/141028096/",
	                "publishedDate" : "Wed, 29 Oct 2014 09:01:38 -0700",
	                "title" : "Butler won't play vs. Knicks; return unclear"
	              },
	              { "author" : "",
	                "categories" : [  ],
	                "content" : "The Bulls open the 2014-15 season on Wednesday in New York with soaring expectations. This is supposed to be the best team of the Tom Thibodeau coaching era, but are the Bulls really better than the 2010-11 version that won 62 games? \"You never know what can happen through the course of a season, but for right now, if we're healthy, why can't we be the best team?\" JImmy Butler said.",
	                "contentSnippet" : "The Bulls open the 2014-15 season on Wednesday in New York with soaring expectations. This is supposed to be the best team of ...",
	                "link" : "http://www.dailyherald.com/article/20141029/sports/141028185/",
	                "publishedDate" : "Wed, 29 Oct 2014 02:38:43 -0700",
	                "title" : "Is this Thibodeau's best Bulls team?"
	              }
	            ],
	          "feedUrl" : "http://www.dailyherald.com/rss/feed/?feed=sports_all",
	          "link" : "http://www.dailyherald.com/sports/",
	          "title" : "DailyHerald.com  > Sports",
	          "type" : "rss20"
	        } },
	  "responseDetails" : null,
	  "responseStatus" : 200
	}
`

func TestFeedResponse(t *testing.T) {
	v := new(FeedResponse)
	err := json.Unmarshal([]byte(feedResponseJSON), v)
	if err != nil {
		t.Errorf("Could not unmarshal json.\nError: %s\n", err.Error())
		return
	}
	if v.ResponseStatus != 200 {
		t.Errorf("Expected response status %d, found %d\n", 200, v.ResponseStatus)
		return
	}
	entries := v.ResponseData.Feed.Entries
	if len(entries) != 2 {
		t.Errorf("Expected %d entries, found %d\n", 2, len(entries))
		return
	}
	entry := entries[0]
	expectedLink := "http://www.dailyherald.com/article/20141029/sports/141028096/"
	if expectedLink != entry.Link {
		t.Errorf("Expected link %s, found %s\n", expectedLink, entry.Link)
		return
	}
}
