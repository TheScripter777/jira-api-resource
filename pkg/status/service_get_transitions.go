package status

import (
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
	"net/http"
)

type ServiceGetTransitions struct {
	issueId string
}

func (s *ServiceGetTransitions) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue

	return service.PreInitJiraAPI(s, params, http.MethodGet)
}

// See service/service.go for details
func (s *ServiceGetTransitions) GetResults() map[string]string {
	return nil
}

// See service/service.go for details
func (s *ServiceGetTransitions) SetResultsFromPrevious(result map[string]string) {
}

// See service/service.go for details
func (s *ServiceGetTransitions) GetEndpoint(url string) string {
	return fmt.Sprintf("%s/issue/%s/transitions", url, s.issueId)
}

// See service/service.go for details
func (s *ServiceGetTransitions) CreateRequestBody() []byte {
	return nil
}

// See service/service.go for details
func (s *ServiceGetTransitions) JSONResponseObject() interface{} {
	return &Transitions{}
}

// See service/service.go for details
func (s *ServiceGetTransitions) PostAPICall(result interface{}) error {
	if transitions, ok := result.(*Transitions); !ok {
		return errors.New("failed to convert result of type interface{} to issue of type reading.Issue")
	} else {
		TransitionsSlice = *transitions
	}
	return nil
}

func (s *ServiceGetTransitions) Name() string {
	return "ServiceGetTransitions"
}

func (s *ServiceGetTransitions) ExecuteAsLastStep(params configuration.JiraAPIResourceParameters) error {
	return nil
}
