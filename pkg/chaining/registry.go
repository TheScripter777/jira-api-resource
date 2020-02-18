package chaining

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/commenting"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/editing"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/noop"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/reading"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/status"
)

const (
	ServiceFetchIssueData      = "srv_fetch_issue_data"
	ServiceReadIssueName       = "srv_read_issue"
	ServiceEditCustomFieldName = "srv_edit_field"
	ServiceGetTransitions      = "srv_get_transitions"
	ServiceDoTransition        = "srv_do_transitions"
	ServiceAddComment          = "srv_add_comment"
	ServiceUnknownName         = "srv_unknown"
)

var serviceRegistry = make(map[string]service.Service)

func InitServiceRegistry() {
	serviceRegistry[ServiceFetchIssueData] = &reading.ServiceFetchIssueData{}
	serviceRegistry[ServiceReadIssueName] = &reading.ServiceReadIssue{}
	serviceRegistry[ServiceEditCustomFieldName] = &editing.ServiceEditCustomField{}
	serviceRegistry[ServiceGetTransitions] = &status.ServiceGetTransitions{}
	serviceRegistry[ServiceDoTransition] = &status.ServiceDoTransition{}
	serviceRegistry[ServiceAddComment] = &commenting.ServiceAddComment{}
	serviceRegistry[ServiceUnknownName] = &noop.ServiceUnknown{}
}

func GetServicesChain(c configuration.Context) []service.Service {
	chain := make([]service.Service, 0)

	switch c {
	case configuration.ReadIssue:
		chain = append(chain, serviceRegistry[ServiceReadIssueName])
	case configuration.ReadStatus:
		chain = append(chain, serviceRegistry[ServiceReadIssueName])
	case configuration.EditCustomField:
		chain = append(chain, serviceRegistry[ServiceReadIssueName])
		chain = append(chain, serviceRegistry[ServiceEditCustomFieldName])
	case configuration.AddComment:
		chain = append(chain, serviceRegistry[ServiceAddComment])
	case configuration.Unknown:
		fallthrough
	default:
		chain = make([]service.Service, 0)
	}

	return chain
}
