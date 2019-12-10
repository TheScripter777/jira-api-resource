package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/log"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/status"
	"net/http"
)

type CreateBodyFN func() []byte
type GetEndpointFN func(string) string
type JsonObjectFN func() interface{}

type JiraAPI struct {
	HttpMethod string
	Body       []byte
	JsonObject interface{}

	// 'url' is not exported because at this level it may contains the credentials
	url string
}

func CreateAPIFromParams(params configuration.JiraAPIResourceParameters, fnBody CreateBodyFN, fnEndpoint GetEndpointFN, fnJsonObj JsonObjectFN, httpMethod string) (JiraAPI, error) {
	var err error
	api := JiraAPI{}

	if fnBody != nil {
		api.Body = fnBody()
	}

	if fnJsonObj != nil {
		api.JsonObject = fnJsonObj()
	}

	api.HttpMethod = httpMethod

	if fnEndpoint == nil {
		return api, errors.New("not allowed to have null endpoint creation function")
	} else {
		api.url = fnEndpoint(*params.JiraAPIUrl)
	}

	return api, err
}

func (api *JiraAPI) Call() (interface{}, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{Transport: tr}
	req, newRequestErr := http.NewRequest(api.HttpMethod, api.url, bytes.NewBuffer(api.Body))

	if newRequestErr != nil {
		return nil, newRequestErr
	}

	if req != nil {
		req.Header.Set("Content-Type", "application/json")

		log.Logger.Debug("Setting http basic auth for api call")
		req.SetBasicAuth(status.Username, status.Password)
	}

	resp, errDo := client.Do(req)

	if errDo != nil {
		return nil, errDo
	}

	defer resp.Body.Close()
	log.Logger.Infof("Received response with %s", fmt.Sprintf("HTTP %s", resp.Status))

	canProcessBody, err := api.processResponse(resp)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	if canProcessBody {
		count, readBodyErr := buffer.ReadFrom(resp.Body)

		//fmt.Printf("Read %v bytes\n", count)
		log.Logger.Debugf("Read %v bytes\n", count)

		if readBodyErr != nil {
			//fmt.Println("Unable to read body of response")
			log.Logger.Error("Unable to read body of response")
			return nil, readBodyErr
		}
		err = json.Unmarshal(buffer.Bytes(), &api.JsonObject)
	}

	return api.JsonObject, err
}

func (api *JiraAPI) processResponse(resp *http.Response) (bool, error) {
	if resp == nil {
		return false, errors.New("nil response was returned")
	}

	if ok, err := Is4xx(resp); ok {
		return false, err
	}

	if ok, err := Is5xx(resp); ok {
		return false, err
	}

	if !HasValidContent(resp) {
		return false, nil
	}

	return true, nil
}
