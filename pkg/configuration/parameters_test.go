// See parameters.go for this package's comment
package configuration

import "testing"

var (
	emptyValue   = ""
	fakeUrl      = "https://github.com/TurnsCoffeeIntoScripts/jira-api-resource"
	fakeUsername = "dummy_username"
	fakePassword = "dummy_password"
)

// Tests the JiraAPIResourceParameters.validate() method
//
func TestValidateSuccessFieldPresentAndValid(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakeUsername,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
		Context:    ReadIssue,
	}

	param.validate()
	testExpectedBoolResult(t, param, param.Meta.valid, true)
}

func TestValidateFailMissingUrl(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &emptyValue,
		Username:   &fakeUsername,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
	}

	param.validate()
	testExpectedBoolResult(t, param, param.Meta.mandatoryPresent, false)
	testExpectedBoolResult(t, param, param.Meta.valid, false)
}

func TestValidateFailMissingUsername(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &emptyValue,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
	}

	param.validate()
	testExpectedBoolResult(t, param, param.Meta.mandatoryPresent, false)
	testExpectedBoolResult(t, param, param.Meta.valid, false)
}

func TestValidateFailMissingPassword(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &emptyValue,
		IssueList:  make([]string, 1),
	}

	param.validate()
	testExpectedBoolResult(t, param, param.Meta.mandatoryPresent, false)
	testExpectedBoolResult(t, param, param.Meta.valid, false)
}

func TestValidateFailMissingIssueList(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &emptyValue,
		Password:   &fakePassword,
		IssueList:  nil,
	}

	param.validate()
	testExpectedBoolResult(t, param, param.Meta.mandatoryPresent, false)
	testExpectedBoolResult(t, param, param.Meta.valid, false)
}

func TestEmptyIssueList(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &fakePassword,
		IssueList:  make([]string, 0),
	}

	param.validate()
	testExpectedBoolResult(t, param, param.Meta.mandatoryPresent, true)
	testExpectedBoolResult(t, param, param.Meta.valid, false)
}

func TestContextUnknown(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
		Context: Unknown,
	}

	param.validate()
	testExpectedBoolResult(t, param, param.Meta.valid, false)
}

func TestContextReadIssue(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
		Context: ReadIssue,
	}

	param.validate()


}

func testExpectedBoolResult(t *testing.T, param *JiraAPIResourceParameters, result, expected bool) {
	// Normally this clause should be written like so: 'if result {'
	// But in the context of this test it makes it easier to read if the '!= expected' is added because
	// the clause can then explicity be read as 'if the result is not expected'
	if result != expected {
		t.Errorf("Boolean value was incorrect, got: %t, want: %t.", result, expected)
	}
}
