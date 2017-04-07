package models

//ReleaseLinks Represents links for a release
type ReleaseLinks struct {
	Self                             string `json:"Self"`
	Project                          string `json:"Project"`
	Progression                      string `json:"Progression"`
	Deployments                      string `json:"Deployments"`
	DeploymentTemplate               string `json:"DeploymentTemplate"`
	Artifacts                        string `json:"Artifacts"`
	ProjectVariableSnapshot          string `json:"ProjectVariableSnapshot"`
	ProjectDeploymentProcessSnapshot string `json:"ProjectDeploymentProcessSnapshot"`
	Web                              string `json:"Web"`
	SnapshotVariables                string `json:"SnapshotVariables"`
	Defects                          string `json:"Defects"`
	ReportDefect                     string `json:"ReportDefect"`
	ResolveDefect                    string `json:"ResolveDefect"`
}
