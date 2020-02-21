package noop

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/http/rest"
)

type ServiceUnknown struct {
}

func (s *ServiceUnknown) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	return rest.JiraAPI{}, nil
}

// See service/service.go for details
func (s *ServiceUnknown) GetResults() map[string]string {
	return nil
}

// See service/service.go for details
func (s *ServiceUnknown) SetResultsFromPrevious(result map[string]string) {
}

// See service/service.go for details
func (s *ServiceUnknown) GetEndpoint(url string) string {
	return fmt.Sprintf("/")
}

// See service/service.go for details
func (s *ServiceUnknown) CreateRequestBody() []byte {
	return nil
}

// See service/service.go for details
func (s *ServiceUnknown) JSONResponseObject() interface{} {
	return nil
}

// See service/service.go for details
func (s *ServiceUnknown) PostAPICall(result interface{}) error {
	return nil
}

func (s *ServiceUnknown) Name() string {
	return "ServiceUnknown"
}

func (s *ServiceUnknown) ExecuteAsLastStep(params configuration.JiraAPIResourceParameters) error {
	return nil
}
