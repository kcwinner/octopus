package models

//DashboardItemLinks Represents the links for a dashboard item
type DashboardItemLinks struct {
	Self    string `json:"Self"`
	Release string `json:"Release"`
	Tenant  string `json:"Tenant"`
	Task    string `json:"Task"`
}
