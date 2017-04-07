package client

import (
	"encoding/json"
	"fmt"

	"github.com/kcwinner/octopus/models"
)

//GetCurrentReleaseForEnvironmentProject Gets a projects release by environment
func (octopus *OctopusClient) GetCurrentReleaseForEnvironmentProject(projectName, environmentName string) (models.Release, models.Project, models.Environment, error) {
	project, err := octopus.GetProjectByName(projectName)
	if err != nil {
		return models.Release{}, models.Project{}, models.Environment{}, err
	}

	environment, err := octopus.GetEnvironment(environmentName)
	if err != nil {
		return models.Release{}, models.Project{}, models.Environment{}, err
	}

	query := fmt.Sprintf("projects=%s&environments=%s", project.ID, environment.ID)

	dashboard, err := octopus.QueryDashboard(query)
	if err != nil {
		return models.Release{}, models.Project{}, models.Environment{}, err
	}

	req, err := octopus.buildOctopusRequest("GET", dashboard.Items[0].Links.Release)
	if err != nil {
		return models.Release{}, models.Project{}, models.Environment{}, err
	}

	resp, err := octopus.client.Do(req)
	if err != nil {
		return models.Release{}, models.Project{}, models.Environment{}, err
	}

	defer resp.Body.Close()

	var release models.Release
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return models.Release{}, models.Project{}, models.Environment{}, err
	}

	return release, project, environment, nil
}
