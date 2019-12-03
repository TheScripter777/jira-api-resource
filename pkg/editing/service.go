// Package editing provides a Jira API interface service and implementation of Jira's domain object as
// Go structures in the context of editing an issue.
package editing

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/helpers"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
	"net/http"
	"strconv"
)

// The ServiceEditCustomField struct implements the service.Service interface. It defines the workflow of editing
// an existing Jira issue.
type ServiceEditCustomField struct {
	issueId    string
	fieldKey   string
	fieldType  string
	fieldValue string
}

// See service/service.go for details
func (s *ServiceEditCustomField) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	s.fieldValue = *params.CustomFieldValue
	s.fieldType = *params.CustomFieldType

	if s.issueId == "" || s.fieldValue == "" || s.fieldType == "" {
		return rest.JiraAPI{}, errors.New("missing value(s) for ServiceEditCustomField")
	}

	return service.PreInitJiraAPI(s, params, http.MethodPut)
}

// See service/service.go for details
func (s *ServiceEditCustomField) GetResults() map[string]string {
	return nil
}

// See service/service.go for details
func (s *ServiceEditCustomField) SetResultsFromPrevious(result map[string]string) {
	s.fieldKey = result[helpers.ReadingFieldKey]
}

// See service/service.go for details
func (s *ServiceEditCustomField) GetEndpoint(url string) string {
	return fmt.Sprintf("%s/issue/%s", url, s.issueId)
}

// See service/service.go for details
func (s *ServiceEditCustomField) CreateRequestBody() []byte {
	i := Issue{}

	if numVal, err := strconv.Atoi(s.fieldValue); err == nil && s.fieldType != "string" {
		i.AddField(s.fieldKey, numVal)
	} else {
		i.AddField(s.fieldKey, s.fieldValue)
	}
	b, err := json.Marshal(i)
	if err != nil {
		b, _ := json.Marshal(Issue{})
		return b
	}
	return b
}

// See service/service.go for details
func (s *ServiceEditCustomField) JSONResponseObject() interface{} {
	return nil
}

// See service/service.go for details
func (s *ServiceEditCustomField) PostAPICall(result interface{}) error {
	return nil
}

func (s *ServiceEditCustomField) Name() string {
	return "ServiceEditCustomField"
}
