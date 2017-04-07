package models

//Deployment Represents and Octopus deployment
type Deployment struct {
	ID                       string            `json:"Id"`
	ChannelID                string            `json:"ChannelId"`
	Comments                 string            `json:"Comments"`
	DeploymentProcessID      string            `json:"DeploymentProcessId"`
	EnvironmentID            string            `json:"EnvironmentId"`
	ForcePackageDownload     bool              `json:"ForcePackageDownload"`
	ForcePackageRedeployment bool              `json:"ForcePackageRedeployment"`
	ProjectID                string            `json:"ProjectId"`
	ReleaseID                string            `json:"ReleaseId"`
	Name                     string            `json:"Name"`
	FormValues               map[string]string `json:"FormValues"`
}
