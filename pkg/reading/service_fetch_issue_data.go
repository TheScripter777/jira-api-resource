package reading

import (
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/helpers"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
	"net/http"
)

type ServiceFetchIssueData struct {
	issueId    string
	parentKey  string
	statusName string
}

func (s *ServiceFetchIssueData) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue

	return service.PreInitJiraAPI(s, params, http.MethodGet)
}

func (s *ServiceFetchIssueData) GetResults() map[string]string {
	var m = make(map[string]string)
	m[helpers.ParentIssueKey] = s.parentKey
	m[helpers.StatusNameKey] = s.statusName
	return m
}

func (s *ServiceFetchIssueData) SetResultsFromPrevious(result map[string]string) {
}

func (s *ServiceFetchIssueData) GetEndpoint(url string) string {
	return fmt.Sprintf("%s/issue/%s?expand=names", url, s.issueId)
}

func (s *ServiceFetchIssueData) CreateRequestBody() []byte {
	return nil
}

func (s *ServiceFetchIssueData) JSONResponseObject() interface{} {
	return &Issue{}
}

func (s *ServiceFetchIssueData) PostAPICall(result interface{}) error {
	if issue, ok := result.(*Issue); !ok {
		return errors.New("failed to convert result of type interface{} to issue of type reading.Issue")
	} else {
		// Find parent key if current one has a parent
		if issue.Fields.Parent != nil {
			s.parentKey = issue.Fields.Parent.Key
		}

		s.statusName = issue.Fields.Status.Name
	}

	return nil
}

func (s *ServiceFetchIssueData) Name() string {
	return "ServiceFetchIssueData"
}

func (s *ServiceFetchIssueData) ExecuteAsLastStep(params configuration.JiraAPIResourceParameters) error {
	return nil
}
