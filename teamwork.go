package teamwork

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	TWHOST = "https://teamwork.com"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientOption func(client *DefaultClient)

type DefaultClient struct {
	uri        string
	token      string
	httpclient httpClient
	log        *log.Logger
}

func NewClient(token string, options ...ClientOption) *DefaultClient {

	c := &DefaultClient{
		uri:        TWHOST,
		token:      token,
		httpclient: &http.Client{},
		log:        log.New(os.Stdout, "teamwork", log.LstdFlags),
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

func AccountPrefixOption(prefix string) ClientOption {

	return func(client *DefaultClient) {

		uri, err := url.Parse(client.uri)
		if err != nil {
			client.log.Fatal("Unable to parse api uri: %v", err)
			return
		}

		hostname := prefix + "." + uri.Hostname()
		if len(uri.Port()) != 0 {
			hostname += ":" + uri.Port()
		}

		uri.Host = hostname
		client.uri = uri.String()
	}
}
