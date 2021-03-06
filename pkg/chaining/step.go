package chaining

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
)

type Step struct {
	Service service.Service
	Name    string
	Last    bool

	params *configuration.JiraAPIResourceParameters
}

func (s *Step) Execute(csValues CrossStepsValues, lastStep bool) error {
	return service.Execute(s.Service, *s.params, lastStep)
}

func (s *Step) PrepareNextStep(ns *Step, csValues CrossStepsValues) CrossStepsValues {
	result := s.Service.GetResults()

	if result != nil {
		if csValues.mapping == nil {
			csValues.mapping = make(map[string]string)
		}

		for k, v := range result {
			csValues.mapping[k] = v
		}

		ns.Service.SetResultsFromPrevious(csValues.mapping)

	}
	return csValues
}
