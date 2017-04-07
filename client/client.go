package client

import (
	"net/http"
	"net/url"
)

//OctopusClient Octopus client struct
type OctopusClient struct {
	Debug  bool
	URL    *url.URL
	APIKey string
	client *http.Client
}

//NewOctopusClient Returns an octopus client
func NewOctopusClient(debug bool, server, apikey string) (*OctopusClient, error) {
	url, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	return &OctopusClient{
		Debug:  debug,
		URL:    url,
		APIKey: apikey,
		client: &http.Client{},
	}, nil
}
