package client

import (
	"encoding/json"
	"fmt"

	"github.com/kcwinner/octopus/models"
)

//GetProjectByName Returns an Octopus project by name
func (octopus *OctopusClient) GetProjectByName(name string) (models.Project, error) {
	path := fmt.Sprintf("/api/projects/%s", name)
	req, err := octopus.buildOctopusRequest("GET", path)
	if err != nil {
		return models.Project{}, err
	}

	resp, err := octopus.client.Do(req)
	if err != nil {
		return models.Project{}, err
	}

	defer resp.Body.Close()

	var project models.Project
	err = json.NewDecoder(resp.Body).Decode(&project)
	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}

//GetProjects Returns all Octopus projects
func (octopus *OctopusClient) GetProjects() ([]models.Project, error) {
	req, err := octopus.buildOctopusRequest("GET", "/api/projects/all")
	if err != nil {
		return []models.Project{}, err
	}

	resp, err := octopus.client.Do(req)
	if err != nil {
		return []models.Project{}, err
	}

	defer resp.Body.Close()

	var projects []models.Project
	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		return []models.Project{}, err
	}

	return projects, nil
}
