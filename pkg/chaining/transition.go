package chaining

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/status"
)

func PerformForceOpen(params *configuration.JiraAPIResourceParameters) bool {
	if err := fetchTransitionsIfMissing(params); err != nil {
		return false
	}

	srvDoTransition := &status.ServiceDoTransition{}
	if err := service.Execute(srvDoTransition, *params, false); err != nil {
		return false
	}

	return true
}

func PerformClose(params *configuration.JiraAPIResourceParameters) error {
	if err := fetchTransitionsIfMissing(params); err != nil {
		return err
	}
	srvDoTransition := &status.ServiceDoTransition{}
	srvDoTransition.OverwriteTransitionName(*params.ClosedStatusName)
	if err := service.Execute(srvDoTransition, *params, false); err != nil {
		return err
	}

	return nil
}

func fetchTransitionsIfMissing(params *configuration.JiraAPIResourceParameters) error {
	if status.TransitionsSlice.Transitions == nil {
		srvGetTransitions := &status.ServiceGetTransitions{}
		if err := service.Execute(srvGetTransitions, *params, false); err != nil {
			return err
		}
	}

	return nil
}
