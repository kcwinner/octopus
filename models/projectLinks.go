package models

//ProjectLinks Represents links to project
type ProjectLinks struct {
	Self              string `json:"Self"`
	Releases          string `json:"Releases"`
	Channels          string `json:"Channels"`
	Triggers          string `json:"Triggers"`
	OrderChannels     string `json:"OrderChannels"`
	Variables         string `json:"Variables"`
	Progression       string `json:"Progression"`
	DeploymentProcess string `json:"DeploymentProcess"`
	Web               string `json:"Web"`
	Logo              string `json:"Logo"`
}
