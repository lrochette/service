package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const v1API = "/api/blog/v1"

// Blog holds client methods
type Blog interface {
	Authors
	Posts
}

type blog struct {
	authorsClient
	postsClient
}

// Config holds client configuration
type Config struct {
	ServiceURI     string
	TimeoutSeconds int
}

type client struct {
	web        *http.Client
	serviceURI string
}

// New inits a blog client
func New(config *Config) Blog {
	web := &http.Client{
		Timeout: time.Second * time.Duration(config.TimeoutSeconds),
	}

	base := client{
		web:        web,
		serviceURI: config.ServiceURI,
	}

	return &blog{
		authorsClient{base},
		postsClient{base},
	}
}

func formatJSONRequest(method string, url string, body interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func executeRequest(client *http.Client, request *http.Request, obj interface{}) (*int, error) {
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if obj == nil {
		return &resp.StatusCode, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, obj)
	return &resp.StatusCode, err
}
