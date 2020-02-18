package chaining

import (
	"errors"
	"fmt"
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
		// If it's the last element, then the lastStep flag is set to true to trigger the output
		p.addStep(chain[s], params, s == len(chain)-1)
	}

	p.length = len(p.steps)

	return nil
}

func (p *Pipeline) addStep(s service.Service, params *configuration.JiraAPIResourceParameters, lastStep bool) {
	newStep := Step{}
	newStep.Service = s
	newStep.Name = s.Name()
	newStep.params = params
	newStep.Last = lastStep

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
	log.Logger.Debug("Executing pipeline for issue ", params.ActiveIssue)
	return p.singleExecution(params)
}

func (p *Pipeline) singleExecution(params *configuration.JiraAPIResourceParameters) error {
	forcedOpen := false

	if err := p.loadIssueData(params); err != nil {
		return err
	}

	for index := range p.steps {
		log.Logger.Debug("Executing step (", index+1, "/", p.length, ")", p.steps[index].Name)
		err := p.steps[index].Execute(p.csValues, p.steps[index].Last)

		if err != nil {
			if helpers.IsBoolPtrTrue(params.Flags.KeepGoingOnError) {
				log.Logger.Warning("Error detected but '--keepGoing' was specified.")
				log.Logger.Warning(err.Error())
				break
			} else {
				return err
			}
		}

		if index < p.length-1 {
			ns := &p.steps[index+1]
			p.csValues = p.steps[index].PrepareNextStep(ns, p.csValues)
		}

		if p.csValues.mapping[helpers.IssueForceOpenKey] != "" {
			if forcedOpen = PerformForceOpen(params); forcedOpen {
				p.csValues.mapping[helpers.IssueForceOpenKey] = ""
			}
		}
	}

	if forcedOpen {
		return PerformClose(params)
	}

	return nil
}

func (p *Pipeline) loadIssueData(params *configuration.JiraAPIResourceParameters) error {
	srvFetchData := &reading.ServiceFetchIssueData{}
	values := CrossStepsValues{}
	values.mapping = make(map[string]string, 0)

	if err := service.Execute(srvFetchData, *params, false); err != nil {
		return err
	}

	results := srvFetchData.GetResults()

	if helpers.IsBoolPtrTrue(params.Flags.ForceOnParent) {
		if results[helpers.ParentIssueKey] != "" {
			log.Logger.Debug(fmt.Sprintf("Setting parent key: %s", results[helpers.ParentIssueKey]))
			params.ActiveIssue = results[helpers.ParentIssueKey]
		}
	}

	if results[helpers.StatusNameKey] == *params.ClosedStatusName {
		if helpers.IsBoolPtrTrue(params.Flags.ForceOpen) {
			values.mapping[helpers.IssueForceOpenKey] = *params.TransitionName
		} else {
			return errors.New(fmt.Sprintf("issue %s is in the '%s' status and the '--forceOpen' flag was not specified", params.ActiveIssue, *params.ClosedStatusName))
		}
	}

	if len(values.mapping) > 0 {
		p.csValues.mapping = helpers.CopyMapString(values.mapping, p.csValues.mapping, true)
	}

	return nil
}
