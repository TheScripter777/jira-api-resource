package chaining

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
)

type Step struct {
	Service service.Service
	Name    string

	params *configuration.JiraAPIResourceParameters
}

func (s *Step) Execute(csValues CrossStepsValues) error {
	return service.Execute(s.Service, *s.params)
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
