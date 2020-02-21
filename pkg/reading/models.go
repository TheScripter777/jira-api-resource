package reading

import "github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/assets"

type VersionReadStatusResponse struct {
	Status string `json:"status"`
}

type VersionReadIssueResponse struct {
	Issues []string `json:"issues"`
}

// TODO to complete. This struct is use to output to a json file as the output of the resource within a concourse pipeline
type InResponseStatus struct {
	Version  VersionReadStatusResponse `json:"version"`
	Metadata assets.Metadata           `json:"metadata"`
}

type InResponseIssue struct {
	Issues   VersionReadIssueResponse `json:"version"`
	Metadata assets.Metadata          `json:"metadata"`
}
