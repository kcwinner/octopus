package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (octopus *OctopusClient) buildOctopusRequest(method, path string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", octopus.URL, path)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Octopus-ApiKey", octopus.APIKey)

	return req, nil
}

func (octopus *OctopusClient) buildOctopusPost(path string, v interface{}) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", octopus.URL, path)

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(v)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Octopus-ApiKey", octopus.APIKey)

	return req, nil
}
