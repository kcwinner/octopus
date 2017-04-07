package models

//Environment Represents and Octopus environment
type Environment struct {
	ID               string           `json:"Id"`
	Name             string           `json:"Name"`
	Description      string           `json:"Description"`
	SortOrder        int              `json:"SortOrder"`
	UseGuidedFailure bool             `json:"UseGuidedFailure"`
	Links            EnvironmentLinks `json:"Links"`
}
