package teamwork

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	TWHOST = "teamwork.com"
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

func NewClient(prefix, token string, options ...ClientOption) *DefaultClient {

	c := &DefaultClient{
		uri:        "https://" + prefix + "." + TWHOST,
		token:      token,
		httpclient: &http.Client{},
		log:        log.New(os.Stdout, "teamwork", log.LstdFlags),
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

func EuUriOption() ClientOption {
	//TODO: Rewrite this function
	return func(client *DefaultClient) {

		uri, err := url.Parse(client.uri)
		if err != nil {
			client.log.Fatalf("Unable to parse api uri: %v", err)
			return
		}

		tmpBuilder := func(parts ...string) string {
			return strings.Join(parts, ".")
		}

		parts := strings.Split(uri.Hostname(), ".")
		build := tmpBuilder(parts[0], "eu", tmpBuilder(parts[1:]...))

		if len(uri.Port()) != 0 {
			build += ":" + uri.Port()
		}

		uri.Host = build
		client.uri = uri.String()
	}
}
