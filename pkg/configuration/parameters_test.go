package configuration

import "testing"

var (
	emptyValue   = ""
	fakeUrl      = "https://github.com/TurnsCoffeeIntoScripts/jira-api-resource"
	fakeUsername = "dummy_username"
	fakePassword = "dummy_password"
)

func TestValidateSuccessFieldPresentAndValid(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakeUsername,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
		Context: ReadIssue,
	}

	param.validate()

	if param.Meta.valid != true {
		t.Errorf("Boolean value was incorrect, got: %t, want: %t.", param.Meta.valid, true)
	}
}

func TestValidateFailMissingUrl(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &emptyValue,
		Username:   &fakeUsername,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
	}

	internalTestMissingValue(t, param)
}

func TestValidateFailMissingUsername(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &emptyValue,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
	}

	internalTestMissingValue(t, param)
}

func TestValidateFailMissingPassword(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &emptyValue,
		IssueList:  make([]string, 1),
	}

	internalTestMissingValue(t, param)
}

func TestValidateFailMissingIssueList(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &emptyValue,
		Password:   &fakePassword,
		IssueList:  nil,
	}

	internalTestMissingValue(t, param)
}

func TestEmptyIssueList(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &fakePassword,
		IssueList:  make([]string, 0),
	}

	param.validate()

	if param.Meta.mandatoryPresent == false {
		t.Errorf("Boolean value was incorrect, got: %t, want: %t.", param.Meta.mandatoryPresent, true)
	}

	if param.Meta.valid != false {
		t.Errorf("Boolean value was incorrect, got: %t, want: %t.", param.Meta.valid, false)
	}
}

func internalTestMissingValue(t *testing.T, param *JiraAPIResourceParameters) {
	param.validate()

	// Normally this clause should be written like so: 'if param.Meta.mandatorypresent {'
	// But in the context of this test it makes it easier to read if the '!= false' is added because
	// the clause can then explicity be read as 'if the flag is not false'
	if param.Meta.mandatoryPresent != false {
		t.Errorf("Boolean value was incorrect, got: %t, want: %t.", param.Meta.mandatoryPresent, false)
	}

	// Normally this clause should be written like so: 'if param.Meta.valid {'
	// But in the context of this test it makes it easier to read if the '!= false' is added because
	// the clause can then explicity be read as 'if the flag is not false'
	if param.Meta.valid != false {
		t.Errorf("Boolean value was incorrect, got: %t, want: %t.", param.Meta.valid, false)
	}
}
