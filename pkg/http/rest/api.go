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
	"strings"
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

	//client := http.DefaultClient
	client := http.Client{Transport: tr}
	req, newRequestErr := http.NewRequest(api.HttpMethod, api.url, bytes.NewBuffer(api.Body))

	if newRequestErr != nil {
		return nil, newRequestErr
	}

	if req != nil {
		req.Header.Set("Content-Type", "application/json")

		if status.Secured {
			log.Logger.Debug("Setting session cookie for api call")
			req.Header.Set("cookie", fmt.Sprintf("%s=%s", status.SessionName, status.SessionValue))
		}
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

// Deprecated: This was used early on in the development to ease the connection
// process. Now the resource uses session cookies.
func (api *JiraAPI) createUrlWithCredentials(url, username, password string) {
	// Construct the user:password@ string
	usrpw := username + ":" + password + "@"

	// Find the index to which we'll insert said string in url
	usrpwInsertIndex := strings.Index(url, "://") + 3

	// Remove any trailing '/' from the url because if needed they'll be added later on
	// by the specific service executing the api call
	strings.TrimSuffix(url, "/")

	// Build the final url value
	api.url = url[:usrpwInsertIndex] + usrpw + url[usrpwInsertIndex:]
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
