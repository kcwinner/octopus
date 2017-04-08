package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kcwinner/octopus/models"
)

func (octopus *OctopusClient) getFormValuesForRelease(releaseID, environmentID string) (map[string]string, error) {
	var formValues map[string]string

	path := fmt.Sprintf("/api/releases/%s/deployments/template", releaseID)
	req, err := octopus.buildOctopusRequest("GET", path)
	if err != nil {
		return formValues, err
	}

	resp, err := octopus.client.Do(req)
	if err != nil {
		return formValues, err
	}

	var template models.DeploymentTemplate
	err = json.NewDecoder(resp.Body).Decode(&template)
	if err != nil {
		return formValues, err
	}

	resp.Body.Close()

	path = template.PromoteTo[0].Links.Preview
	req, err = octopus.buildOctopusRequest("GET", path)
	if err != nil {
		log.Println(err)
		return formValues, err
	}

	resp, err = octopus.client.Do(req)
	if err != nil {
		return formValues, err
	}

	var preview models.DeploymentPreview
	err = json.NewDecoder(resp.Body).Decode(&preview)
	if err != nil {
		return formValues, err
	}

	resp.Body.Close()

	formValues = make(map[string]string)
	for _, element := range preview.Form.Elements {
		formValues[element.Control.Name] = element.Name
	}

	return formValues, nil
}

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
