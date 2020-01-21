package reading

import "github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/assets"

type VersionStatus struct {
	Status string `json:"status"`
}

type InResponse struct {
	Version  VersionStatus   `json:"version"`
	Metadata assets.Metadata `json:"metadata"`
}
