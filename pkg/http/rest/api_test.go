package rest

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"net/http"
	"testing"
)

var (
	emptyValue           = ""
	fakeUrl              = "https://github.com/TurnsCoffeeIntoScripts/jira-api-resource"
	fakeUsername         = "dummy_username"
	fakePassword         = "dummy_password"
	fakeCustomFieldValue = "dummyValue"
)

type JiraAPIMock struct {
}

type DummyJsonObject struct {
}

func (api *JiraAPIMock) Call() (interface{}, error) {
	return nil, nil
}

func (api *JiraAPIMock) processResponse(resp *http.Response) (bool, error) {
	return false, nil
}

func TestCreateAPIFromParamsSuccess(t *testing.T) {
	param := configuration.JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
	}

	api, err := CreateAPIFromParams(param,
		createEmptyBody,
		getDummyEndpoint,
		jsonObject,
		http.MethodGet)

	if err != nil {
		t.Errorf("Error value was incorrect, got: %v, want: %v.", err, nil)
	}

	if api.url == "" {
		t.Errorf("String value was incorrect, got: an empty string, want: any non-empty string.")
	}
}

func TestCreateApiFromParamsEmptyBodySuccess(t *testing.T) {
	param := configuration.JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
	}

	api, err := CreateAPIFromParams(param,
		nil,
		getDummyEndpoint,
		jsonObject,
		http.MethodGet)

	if err != nil {
		t.Errorf("Error value was incorrect, got: %v, want: %v.", err, nil)
	}

	if api.Body != nil {
		t.Errorf("[]byte (body) value was incorrect, got: %v, want: %v.", api.Body, nil)
	}
}

func createEmptyBody() []byte {
	return make([]byte, 0)
}

func getDummyEndpoint(s string) string {
	return "dummy/" + s
}

func jsonObject() interface{} {
	return &DummyJsonObject{}
}
