// See parameters.go for this package's comment
package configuration

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var (
	emptyValue           = ""
	fakeUrl              = "https://github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource"
	fakeUsername         = "dummy_username"
	fakePassword         = "dummy_password"
	fakeCustomFieldValue = "dummyValue"
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
	testExpectedBoolResult(t, param.Meta.valid, true)
}

func TestValidateFailMissingUrl(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &emptyValue,
		Username:   &fakeUsername,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.mandatoryPresent, false)
	testExpectedBoolResult(t, param.Meta.valid, false)
}

func TestValidateFailMissingUsername(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &emptyValue,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.mandatoryPresent, false)
	testExpectedBoolResult(t, param.Meta.valid, false)
}

func TestValidateFailMissingPassword(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &emptyValue,
		IssueList:  make([]string, 1),
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.mandatoryPresent, false)
	testExpectedBoolResult(t, param.Meta.valid, false)
}

func TestValidateFailMissingIssueList(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &emptyValue,
		Password:   &fakePassword,
		IssueList:  nil,
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.mandatoryPresent, false)
	testExpectedBoolResult(t, param.Meta.valid, false)
}

func TestEmptyIssueList(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &fakePassword,
		IssueList:  make([]string, 0),
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.mandatoryPresent, true)
	testExpectedBoolResult(t, param.Meta.valid, false)
}

func TestContextUnknown(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
		Context:    Unknown,
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.valid, false)
}

func TestContextReadIssue(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
		Context:    ReadIssue,
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.valid, true)
}

func TestContextEditCustomFieldSuccess1(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl:       &fakeUrl,
		Username:         &fakePassword,
		Password:         &fakePassword,
		IssueList:        make([]string, 1),
		Context:          EditCustomField,
		CustomFieldValue: &fakeCustomFieldValue,
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.valid, true)
}

func TestContextEditCustomFieldSuccess2(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl:               &fakeUrl,
		Username:                 &fakePassword,
		Password:                 &fakePassword,
		IssueList:                make([]string, 1),
		Context:                  EditCustomField,
		CustomFieldValueFromFile: &fakeCustomFieldValue,
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.valid, true)
}

func TestContextEditCustomFieldFailMissingBothValuesNil(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl: &fakeUrl,
		Username:   &fakePassword,
		Password:   &fakePassword,
		IssueList:  make([]string, 1),
		Context:    EditCustomField,
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.valid, false)

	if param.Meta.Msg == "" {
		t.Errorf("String value was incorrect, got empty string but expected a message")
	}
}

func TestContextEditCustomFieldFailMissingBothValuesEmpty(t *testing.T) {
	param := &JiraAPIResourceParameters{
		JiraAPIUrl:               &fakeUrl,
		Username:                 &fakePassword,
		Password:                 &fakePassword,
		IssueList:                make([]string, 1),
		Context:                  EditCustomField,
		CustomFieldValue:         &emptyValue,
		CustomFieldValueFromFile: &emptyValue,
	}

	param.validate()
	testExpectedBoolResult(t, param.Meta.valid, false)

	if param.Meta.Msg == "" {
		t.Errorf("String value was incorrect, got empty string but expected a message")
	}
}

func TestInitializeIssueListSingle1(t *testing.T) {
	param := &JiraAPIResourceParameters{}

	list := "ABC-001"
	param.initializeIssueList(&list)

	if param.IssueList == nil {
		t.Errorf("IssueList was nil but should have contained at least one element")
	}

	testExpectedBoolResult(t, param.Meta.MultipleIssue, false)
}

func TestInitializeIssueListSingle2(t *testing.T) {
	param := &JiraAPIResourceParameters{}

	list := "ABC-001,"
	param.initializeIssueList(&list)

	if param.IssueList == nil {
		t.Errorf("IssueList was nil but should have contained at least one element")
	}

	testExpectedBoolResult(t, param.Meta.MultipleIssue, false)
}

func TestInitializeIssueListMultiple(t *testing.T) {
	param := &JiraAPIResourceParameters{}

	list := "ABC-001,DEF-9999"
	param.initializeIssueList(&list)

	if param.IssueList == nil {
		t.Errorf("IssueList was nil but should have contained at least one element")
	}

	testExpectedBoolResult(t, param.Meta.MultipleIssue, true)
}

func TestInitializeContext(t *testing.T) {
	dummyString := "abc"
	readIssueString := "ReadIssue"
	editCustomFieldString := "EditCustomField"

	tables := []struct {
		i *string
		o Context
	}{
		{nil, Unknown},
		{&dummyString, Unknown},
		{&readIssueString, ReadIssue},
		{&editCustomFieldString, EditCustomField},
	}

	for _, table := range tables {
		param := &JiraAPIResourceParameters{}
		param.initializeContext(table.i)

		if param.Context != table.o {
			t.Errorf("Context value was incorrect, got: %v, want: %v.", param.Context, table.o)
		}
	}
}

func TestParsePtrParameters(t *testing.T) {
	tables := []struct {
		in1  string
		in2  string
		out1 string
	}{
		{"url", "github.com", "JiraAPIUrl"},
		{"username", "I_AM_USER", "Username"},
		{"password", "SECRET_PWD", "Password"},
		{"customFieldName", "Custom_Field_Name", "CustomFieldName"},
		{"customFieldType", "String", "CustomFieldType"},
		{"customFieldValue", "123abc", "CustomFieldValue"},
		{"customFieldValueFromFile", "filename.dummy", "CustomFieldValueFromFile"},
		{"loggingLevel", "TEST", "LoggingLevel"},
	}

	for _, table := range tables {
		param := &JiraAPIResourceParameters{}
		runFlagTest(t, param, table.in1, table.in2, table.out1)
	}
}

func runFlagTest(t *testing.T, param *JiraAPIResourceParameters, p, expected string, refVal string) {
	os.Args = setupOsArgs(fmt.Sprintf("--%s=%s", p, expected))

	param.Parse()

	rParam := reflect.Indirect(reflect.ValueOf(param))
	f := rParam.FieldByName(refVal)

	testExpectedStringResult(t, *f.Interface().(*string), expected)

	// Reset for next test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

}

func setupOsArgs(vals ...string) []string {
	args := make([]string, 0)
	args = append(args, "parameters_test")
	for _, val := range vals {
		args = append(args, val)
	}

	return args
}

func testExpectedStringResult(t *testing.T, result, expected string) {
	// Normally this clause should be written like so: 'if result {'
	// But in the context of this test it makes it easier to read if the '!= expected' is added because
	// the clause can then explicity be read as 'if the result is not expected'
	if result != expected {
		t.Errorf("String value was incorrect, got: %s, want: %s.", result, expected)
	}
}

func testExpectedBoolResult(t *testing.T, result, expected bool) {
	// Normally this clause should be written like so: 'if result {'
	// But in the context of this test it makes it easier to read if the '!= expected' is added because
	// the clause can then explicity be read as 'if the result is not expected'
	if result != expected {
		t.Errorf("Boolean value was incorrect, got: %t, want: %t.", result, expected)
	}
}
