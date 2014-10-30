package feeder

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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

func (a *API) feedLoadEndpoint(feedUrl string) (*url.URL, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("v", string(a.Version))
	q.Set("num", strconv.Itoa(int(a.Number)))
	q.Set("q", feedUrl)
	u.RawQuery = q.Encode()
	return u, nil
}

func (a *API) FeedLoad(feedUrl string) (*FeedResponse, error) {
	url, err := a.feedLoadEndpoint(feedUrl)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url.String())
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
