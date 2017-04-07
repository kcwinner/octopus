package models

//DashboardItem Represents an item on the Octopus Dashboard
type DashboardItem struct {
	ID             string             `json:"Id"`
	ProjectID      string             `json:"ProjectId"`
	EnvironmentID  string             `json:"EnvironmentId"`
	ReleaseID      string             `json:"ReleaseId"`
	DeploymentID   string             `json:"DeploymentId"`
	TaskID         string             `json:"TaskId"`
	ChannelID      string             `json:"ChannelId"`
	ReleaseVersion string             `json:"ReleaseVersion"`
	Created        string             `json:"Created"`
	State          string             `json:"State"`
	ErrorMessage   string             `json:"ErrorMessage"`
	Duration       string             `json:"Duration"`
	IsCurrent      bool               `json:"IsCurrent"`
	IsPrevious     bool               `json:"IsPrevious"`
	IsCompleted    bool               `json:"IsCompleted"`
	Links          DashboardItemLinks `json:"Links"`
}
