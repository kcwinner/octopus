package models

//DeploymentTemplate Represents an Octopus deployment template
type DeploymentTemplate struct {
	IsDeploymentProcessModified  bool                    `json:"IsDeploymentProcessModified"`
	IsVariableSetModified        bool                    `json:"IsVariableSetModified"`
	IsLibraryVariableSetModified bool                    `json:"IsLibraryVariableSetModified"`
	PromoteTo                    []PromoteTo             `json:"PromoteTo"`
	Links                        DeploymentTemplateLinks `json:"Links"`
}

//DeploymentTemplateLinks Octopus deployment template links
type DeploymentTemplateLinks struct {
	Self    string `json:"Self"`
	Release string `json:"Release"`
}

//PromoteTo Reprensents an environment that can be promoted to
type PromoteTo struct {
	ID    string         `json:"Id"`
	Name  string         `json:"Name"`
	Links PromoteToLinks `json:"Links"`
}

//PromoteToLinks Links for promotion
type PromoteToLinks struct {
	Self    string `json:"Self"`
	Preview string `json:"Preview"`
}
