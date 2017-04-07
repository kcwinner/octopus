package models

//Dashboard Represents the Octopus dashboard view
type Dashboard struct {
	Projects     []Project       `json:"Projects"`
	Environments []Environment   `json:"Environments"`
	Items        []DashboardItem `json:"Items"`
	IsFiltered   bool            `json:"IsFiltered"`
}
