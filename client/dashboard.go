package client

import (
	"encoding/json"
	"fmt"

	"github.com/kcwinner/octopus/models"
)

//QueryDashboard Queries the dashboard
func (octopus *OctopusClient) QueryDashboard(query string) (models.Dashboard, error) {
	path := fmt.Sprintf("/api/dashboard/dynamic?%s", query)
	req, err := octopus.buildOctopusRequest("GET", path)
	if err != nil {
		return models.Dashboard{}, err
	}

	resp, err := octopus.client.Do(req)
	if err != nil {
		return models.Dashboard{}, err
	}

	defer resp.Body.Close()

	var result models.Dashboard
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return models.Dashboard{}, err
	}

	return result, nil
}
