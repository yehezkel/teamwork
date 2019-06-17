package teamwork

import (
	"log"
	"net/http"
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



