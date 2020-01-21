package chaining

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/editing"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/noop"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/reading"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
)

const (
	ServiceReadIssueName       = "srv_read_issue"
	ServiceEditCustomFieldName = "srv_edit_field"
	ServiceUnknownName         = "srv_unknown"
)

var serviceRegistry = make(map[string]service.Service)

func InitServiceRegistry() {
	serviceRegistry[ServiceReadIssueName] = &reading.ServiceReadIssue{}
	serviceRegistry[ServiceEditCustomFieldName] = &editing.ServiceEditCustomField{}
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
	case configuration.Unknown:
		fallthrough
	default:
		chain = make([]service.Service, 0)
	}

	return chain
}
