package chaining

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/helpers"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/log"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/reading"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
)

type Pipeline struct {
	steps    []Step
	length   int
	csValues CrossStepsValues
}

func (p *Pipeline) BuildPipelineFromChain(chain []service.Service, params *configuration.JiraAPIResourceParameters) error {
	for s := range chain {
		p.addStep(chain[s], params)
	}

	p.length = len(p.steps)

	return nil
}

func (p *Pipeline) addStep(s service.Service, params *configuration.JiraAPIResourceParameters) {
	newStep := Step{}
	newStep.Service = s
	newStep.Name = s.Name()
	newStep.params = params

	p.steps = append(p.steps, newStep)
}

func (p *Pipeline) Execute(params *configuration.JiraAPIResourceParameters) error {
	if params.Meta.MultipleIssue {
		for _, i := range params.IssueList {
			params.ActiveIssue = i
			log.Logger.Debug("Executing pipeline for issue ", i)
			err := p.singleExecution(params)

			if err != nil {
				return err
			}
		}
	}

	params.ActiveIssue = params.IssueList[0]
	return p.singleExecution(params)
}

func (p *Pipeline) singleExecution(params *configuration.JiraAPIResourceParameters) error {
	for index := range p.steps {
		if params.Flags.ForceOnParent != nil && *params.Flags.ForceOnParent {
			log.Logger.Debug("Executing pre-step to validate parent status")
			tempService := &reading.ServiceReadIssue{}
			tempErr := service.Execute(tempService, *params)
			if tempErr != nil {
				return tempErr
			}

			if tempService.GetResults()[helpers.ParentIssueKey] != "" {
				params.ActiveIssue = tempService.GetResults()[helpers.ParentIssueKey]
			}
		}

		log.Logger.Debug("Executing step (", index+1, "/", p.length, ")", p.steps[index].Name)
		err := p.steps[index].Execute(p.csValues)

		if err != nil {
			return err
		}

		if index < p.length-1 {
			ns := &p.steps[index+1]
			p.csValues = p.steps[index].PrepareNextStep(ns, p.csValues)
		}
	}

	return nil
}
