package models

//Project Represents an Octopus project
type Project struct {
	ID                              string       `json:"Id"`
	VariableSetID                   string       `json:"VariableSetId"`
	DeploymentProcessID             string       `json:"DeploymentProcessId"`
	IncludedLibraryVariableSetIDs   []string     `json:"IncludedLibraryVariableSetIds"`
	DefaultToSkipIfAlreadyInstalled bool         `json:"DefaultToSkipIfAlreadyInstalled"`
	TenantedDeploymentMode          string       `json:"TenantedDeploymentMode"`
	Name                            string       `json:"Name"`
	Slug                            string       `json:"Slug"`
	Description                     string       `json:"Description"`
	IsDisabled                      bool         `json:"IsDisabled"`
	ProjectGroupID                  string       `json:"ProjectGroupId"`
	LifecycleID                     string       `json:"LifecycleId"`
	AutoCreateRelease               bool         `json:"AutoCreateRelease"`
	Links                           ProjectLinks `json:"Links"`
}
