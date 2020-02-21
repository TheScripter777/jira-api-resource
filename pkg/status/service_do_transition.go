package status

import (
	"encoding/json"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
	"net/http"
)

type ServiceDoTransition struct {
	issueId    string
	statusName string
}

func (s *ServiceDoTransition) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	if s.statusName == "" {
		s.statusName = *params.TransitionName
	}

	return service.PreInitJiraAPI(s, params, http.MethodPost)
}

// See service/service.go for details
func (s *ServiceDoTransition) GetResults() map[string]string {
	var m = make(map[string]string)

	return m
}

// See service/service.go for details
func (s *ServiceDoTransition) SetResultsFromPrevious(result map[string]string) {
}

// See service/service.go for details
func (s *ServiceDoTransition) GetEndpoint(url string) string {
	return fmt.Sprintf("%s/issue/%s/transitions", url, s.issueId)
}

// See service/service.go for details
func (s *ServiceDoTransition) CreateRequestBody() []byte {
	t := DoTransitionObject{}

	var tId string
	// Lookup id from stored status name
	for _, tVal := range TransitionsSlice.Transitions {
		if tVal.To.Name == s.statusName {
			tId = tVal.Id
		}
	}

	t.Transition.Id = tId

	b, err := json.Marshal(t)
	if err != nil {
		b, _ := json.Marshal(InnerTransition{})
		return b
	}
	return b
}

// See service/service.go for details
func (s *ServiceDoTransition) JSONResponseObject() interface{} {
	return nil
}

// See service/service.go for details
func (s *ServiceDoTransition) PostAPICall(result interface{}) error {
	return nil
}

func (s *ServiceDoTransition) Name() string {
	return "ServiceDoTransition"
}

func (s *ServiceDoTransition) ExecuteAsLastStep(params configuration.JiraAPIResourceParameters) error {
	return nil
}

func (s *ServiceDoTransition) OverwriteTransitionName(name string) {
	s.statusName = name
}
