package models

//DeploymentPreview Octopus deployment preview
type DeploymentPreview struct {
	StepsToExecute                []DeploymentPreviewStepsToExecute `json:"StepsToExecute"`
	Form                          DeploymentPreviewForm             `json:"Form"`
	UseGuidedFailureModeByDefault bool                              `json:"UseGuidedFailureModeByDefault"`
}

//DeploymentPreviewStepsToExecute Steps to be execute by deployment
type DeploymentPreviewStepsToExecute struct {
	ActionID                string   `json:"ActionId"`
	ActionName              string   `json:"ActionName"`
	ActionNumber            string   `json:"ActionNumber"`
	Roles                   []string `json:"Roles"`
	MachineNames            []string `json:"MachineNames"`
	CanBeSkipped            bool     `json:"canBeSkipped"`
	HasNoApplicableMachines bool     `json:"HasNoApplicableMachines"`
}

//DeploymentPreviewForm Form for Deployment Preview
type DeploymentPreviewForm struct {
	Values   interface{}                    `json:"Values"`
	Elements []DeploymentPreviewFormElement `json:"Elements"`
}

//DeploymentPreviewFormElement Element for Octopus Deployment Preview Form
type DeploymentPreviewFormElement struct {
	Name            string                              `json:"Name"`
	Control         DeploymentPreviewFormElementControl `json:"Control"`
	IsValueRequired bool                                `json:"IsValueRequired"`
}

//DeploymentPreviewFormElementControl Control for Octopus Deployment Preview Form Element
type DeploymentPreviewFormElementControl struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	Label       string `json:"Label"`
	Description string `json:"Description"`
	IsSecure    bool   `json:"IsSecure"`
}
