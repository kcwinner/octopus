package models

//Release Represents an Octopus release
type Release struct {
	ID                           string       `json:"Id"`
	Assembled                    string       `json:"Assembled"`
	ProjectID                    string       `json:"ProjectId"`
	ChannelID                    string       `json:"ChannelId"`
	ProjectVariableSetSnapshotID string       `json:"ProjectVariableSetSnapshotId"`
	Version                      string       `json:"Version"`
	Links                        ReleaseLinks `json:"Links"`
}
