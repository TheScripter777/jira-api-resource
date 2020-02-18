package reading

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/helpers"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/result"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
	"net/http"
	"os"
)

type ServiceReadIssue struct {
	SkipCustomKeyRetrieval bool

	issueId     string
	fieldKey    string
	fieldName   string
	statusName  string
	destination string
}

func (s *ServiceReadIssue) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	s.fieldName = *params.EditCustomFieldParam.CustomFieldName
	s.destination = *params.Destination

	return service.PreInitJiraAPI(s, params, http.MethodGet)
}

func (s *ServiceReadIssue) GetResults() map[string]string {
	var m = make(map[string]string)
	m[helpers.ReadingFieldKey] = s.fieldKey
	return m
}

func (s *ServiceReadIssue) SetResultsFromPrevious(result map[string]string) {
}

func (s *ServiceReadIssue) GetEndpoint(url string) string {
	return fmt.Sprintf("%s/issue/%s?expand=names", url, s.issueId)
}

func (s *ServiceReadIssue) CreateRequestBody() []byte {
	return nil
}

func (s *ServiceReadIssue) JSONResponseObject() interface{} {
	return &Issue{}
}

func (s *ServiceReadIssue) PostAPICall(result interface{}) error {
	if issue, ok := result.(*Issue); !ok {
		return errors.New("failed to convert result of type interface{} to issue of type reading.Issue")
	} else {
		// Match custom field name if it was set
		if !s.SkipCustomKeyRetrieval && s.fieldName != "" {
			s.fieldKey = helpers.FindCustomName(issue.Names.CustomFields, s.fieldName)

			if s.fieldKey == "" {
				return errors.New("failed to retrieve field key from specified custom field name")
			}
		}

		// Find the id of the status of the current Jira issue
		if issue.Fields.Status != nil {
			s.statusName = issue.Fields.Status.Name
		}
	}

	return nil
}

func (s *ServiceReadIssue) Name() string {
	return "ServiceReadIssue"
}

func (s *ServiceReadIssue) ExecuteAsLastStep(ctx configuration.Context) error {
	if file, err := result.CreateDestination(s.destination); err != nil {
		return err
	} else {

		switch ctx {
		case configuration.ReadStatus:
			vs := VersionStatus{Status: s.statusName}
			err := json.NewEncoder(os.Stdout).Encode(InResponse{
				Version: vs,
			})

			if err != nil {
				return err
			}

			return result.Write(file, s.statusName)
		}
	}

	return nil
}
