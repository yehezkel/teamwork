package teamwork

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type ApiClient interface {
	DoRequest(method, path string, payload, out interface{}) error
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
		token:      EncodeApiKey(token),
		httpclient: &http.Client{},
		log:        log.New(os.Stdout, "teamwork", log.LstdFlags),
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

func (client *DefaultClient) DoRequest(method, path string, payload, out interface{}) error {

	uri := client.uri + "/" + path
	var err error
	var request *http.Request

	if method == "GET" {
		request, err = http.NewRequest(method, uri, nil)
	} else {

		//json encode payload
		b, err := json.Marshal(payload)
		//client.log.Printf("uri: %s Payload: %v", uri, string(b))
		if err != nil {
			return err
		}

		request, err = http.NewRequest(
			method, uri, bytes.NewReader(b),
		)
	}

	if err != nil {
		return err
	}
	//add payload for post,put,del
	request.Header.Add("Authorization", "Basic "+client.token)

	resp, err := client.httpclient.Do(request)
	if err != nil {
		return err
	}

	//client.log.Printf("raw response: %#v", resp)

	if (resp.StatusCode < 200 || resp.StatusCode > 299) && resp.StatusCode != 304 {
		return fmt.Errorf("Unexpected response code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	//client.log.Printf("Response: %v", string(body))
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	//client.log.Printf("Response: %v", out)

	return nil
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

func EncodeApiKey(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}
