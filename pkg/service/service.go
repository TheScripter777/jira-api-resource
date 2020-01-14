package service

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/http/rest"
)

type Service interface {
	// Method that will init and return a rest.JiraAPI struct. The initialization is done according
	// to a specific context
	InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error)

	// This method will return a map containing the result(s) of execution of the service. This method is used only
	// when the service is used as a "pre-service" (called before the main service of the specified context)
	GetResults() map[string]string

	// This method takes a map as input that represents the result(s) of a previous service. Each service will fetch
	// and use the value of this map as it sees fit.
	SetResultsFromPrevious(results map[string]string)

	// Returns the parametrized endpoint (what comes after the /rest/api/).
	GetEndpoint(url string) string

	// Returns a slice of byte representing the JSON body of the request (if any).
	CreateRequestBody() []byte

	// Returns a new object (struct) that is the implementation of the Jira object.
	JSONResponseObject() interface{}

	// Any custom operation to be performed after the API call. Such as setting value in the map
	// returned by GetResults()
	PostAPICall(result interface{}) error

	Name() string
}

func PreInitJiraAPI(s Service, params configuration.JiraAPIResourceParameters, httpMethod string) (rest.JiraAPI, error) {
	api, err := rest.CreateAPIFromParams(params, s.CreateRequestBody, s.GetEndpoint, s.JSONResponseObject, httpMethod)
	if err != nil {
		return api, err
	}

	return api, nil
}

func Execute(s Service, params configuration.JiraAPIResourceParameters) error {
	result, err := exec(s, params)

	if err != nil {
		return err
	}

	return s.PostAPICall(result)
}

func exec(s Service, params configuration.JiraAPIResourceParameters) (interface{}, error) {
	api, err := s.InitJiraAPI(params)

	if err != nil {
		return nil, err
	} else {
		return api.Call()
	}
}
