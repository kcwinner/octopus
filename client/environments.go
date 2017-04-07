package client

import (
	"encoding/json"
	"errors"

	"github.com/kcwinner/octopus/models"
)

//GetEnvironment Gets an Octopus environment by name
func (octopus *OctopusClient) GetEnvironment(name string) (models.Environment, error) {
	environments, err := octopus.GetEnvironments()
	if err != nil {
		return models.Environment{}, err
	}

	for _, environment := range environments {
		if environment.Name == name {
			return environment, nil
		}
	}

	return models.Environment{}, errors.New("Environment does not exist")
}

//GetEnvironments Returns all Octopus environments
func (octopus *OctopusClient) GetEnvironments() ([]models.Environment, error) {
	req, err := octopus.buildOctopusRequest("GET", "/api/environments/all")
	if err != nil {
		return nil, err
	}

	resp, err := octopus.client.Do(req)
	if err != nil {
		return []models.Environment{}, err
	}

	defer resp.Body.Close()

	var environments []models.Environment
	err = json.NewDecoder(resp.Body).Decode(&environments)
	if err != nil {
		return []models.Environment{}, err
	}

	return environments, nil
}
