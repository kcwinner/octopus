package client

import (
	"io/ioutil"
	"log"

	"github.com/kcwinner/octopus/models"
)

//PromoteRelease Promotes a release from one environment to the next
func (octopus *OctopusClient) PromoteRelease(projectName, from, to string, formValues map[string]string) error {
	if octopus.Debug {
		log.Println("Getting current release for Environment and project")
	}

	release, project, _, err := octopus.GetCurrentReleaseForEnvironmentProject(projectName, from)
	if err != nil {
		return err
	}

	environment, err := octopus.GetEnvironment(to)
	if err != nil {
		return err
	}

	if len(formValues) > 0 {
		namedValues, err := octopus.getFormValuesForRelease(release.ID, environment.ID)
		if err != nil {
			return err
		}

		deployVals := make(map[string]string)
		for key, val := range namedValues {
			_, ok := formValues[key]
			if ok {
				deployVals[val] = formValues[key]
			}
		}

		formValues = deployVals
	}

	deployment := models.Deployment{
		Comments:                 "Deployed from CLI",
		EnvironmentID:            environment.ID,
		ForcePackageDownload:     false,
		ForcePackageRedeployment: false,
		ProjectID:                project.ID,
		ReleaseID:                release.ID,
		FormValues:               formValues,
	}

	req, err := octopus.buildOctopusPost("/api/deployments", deployment)
	if err != nil {
		return err
	}

	log.Println(req)

	resp, err := octopus.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))

	return nil
}
