package commenting

import (
	"encoding/json"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
	"net/http"
)

type ServiceAddComment struct {
	issueId     string
	commentBody string
}

func (s *ServiceAddComment) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	s.commentBody = *params.AddComment.CommentBody

	return service.PreInitJiraAPI(s, params, http.MethodPost)
}

// See service/service.go for details
func (s *ServiceAddComment) GetResults() map[string]string {
	return nil
}

// See service/service.go for details
func (s *ServiceAddComment) SetResultsFromPrevious(result map[string]string) {
}

// See service/service.go for details
func (s *ServiceAddComment) GetEndpoint(url string) string {
	return fmt.Sprintf("%s/issue/%s/comment", url, s.issueId)
}

// See service/service.go for details
func (s *ServiceAddComment) CreateRequestBody() []byte {
	c := Comment{Body: s.commentBody}

	b, err := json.Marshal(c)
	if err != nil {
		b, _ := json.Marshal(Comment{})
		return b
	}

	return b
}

// See service/service.go for details
func (s *ServiceAddComment) JSONResponseObject() interface{} {
	return nil
}

// See service/service.go for details
func (s *ServiceAddComment) PostAPICall(result interface{}) error {
	return nil
}

func (s *ServiceAddComment) Name() string {
	return "ServiceAddComment"
}

func (s *ServiceAddComment) ExecuteAsLastStep(ctx configuration.Context) error {
	return nil
}
