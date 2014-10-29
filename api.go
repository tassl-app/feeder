package feeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIVersion string

const (
	APIVersion1 = APIVersion("1.0")
	APIVersion2 = APIVersion("2.0")
)

const endpoint = "https://ajax.googleapis.com/ajax/services/feed/load"

type API struct {
	Version APIVersion
	Number  int64
}

func NewAPI() *API {
	return &API{
		Version: APIVersion2,
		Number:  25,
	}
}

func (a *API) feedLoadEndpoint(feedUrl string) string {
	return fmt.Sprintf("%s?v=%s&num=%d&q=%s", endpoint, a.Version, a.Number, feedUrl)
}

func (a *API) FeedLoad(feedUrl string) (*FeedResponse, error) {
	endpoint := a.feedLoadEndpoint(feedUrl)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	feedResponse := new(FeedResponse)
	err = json.Unmarshal(body, feedResponse)
	return feedResponse, err
}
