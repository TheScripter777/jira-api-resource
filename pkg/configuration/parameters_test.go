package configuration_test

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestJiraAPIResourceParameters struct {
	JiraAPIUrl       string   `faker:"url"`
	Username         string   `faker:"username"`
	Password         string   `faker:"password"`
	Destination      string   `faker:"word"`
	IssueList        []string `faker:"len=5"`
	ClosedStatusName string   `faker:"word"`
	TransitionName   string   `faker:"word"`
	Context          configuration.Context

	EditCustomFieldParam TestJiraApiResourceParametersEditCustomField
	AddComment           JiraApiResourceParametersAddComment

	ActiveIssue string
}

type TestJiraApiResourceParametersEditCustomField struct {
	CustomFieldName          string `faker:"word"`
	CustomFieldType          string `faker:"word"`
	CustomFieldValue         string `faker:"word"`
	CustomFieldValueFromFile string `faker:"word"`
}

type JiraApiResourceParametersAddComment struct {
	CommentBody *string `faker:"sentence"`
}

var tParam TestJiraAPIResourceParameters
var context string
var issueList string
var fPtr *bool
var tPtr *bool

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test cases...")
	tParam = TestJiraAPIResourceParameters{}

	err := faker.FakeData(&tParam)
	require.NoError(t, err, "problem creating fake data")

	tParam.Destination = tParam.Destination + "/"

	fPtr = new(bool)
	*fPtr = false
	tPtr = new(bool)
	*tPtr = true

	return func(t *testing.T) {
		t.Log("teardown test cases...")
	}
}

func TestPostParse_Validate(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	param := configuration.JiraAPIResourceParameters{}
	_, _ = param.Parse()

	t.Run("app parameters VALID AND READY from VALID inputs", func(t *testing.T) {
		// Arrange
		param = convertToJiraApiResourceParameters(param)
		context = "ReadIssue"
		issueList = "ABC-123 DEF-456"

		// Act
		param.InitializeAndValidatePostParse(&context, &issueList)

		// Assert
		assert.True(t, param.Meta.AllMandatoryValuesPresent(), "method AllMandatoryValuesPresent() returned false")
		assert.True(t, param.Meta.Ready(), "method Ready() returned false")
	})

	t.Run("app parameters NOT VALID NOR READY from INVALID inputs (MISSING URL)", func(t *testing.T) {
		// Arrange
		param = convertToJiraApiResourceParameters(param)
		*param.JiraAPIUrl = ""
		context = "ReadIssue"
		issueList = "ABC-123 DEF-456"

		// Act
		param.InitializeAndValidatePostParse(&context, &issueList)

		// Assert
		assert.False(t, param.Meta.AllMandatoryValuesPresent(), "method AllMandatoryValuesPresent() returned true")
		assert.False(t, param.Meta.Ready(), "method Ready() returned true")
	})

	t.Run("app parameters NOT VALID NOR READY from INVALID inputs (MISSING USERNAME)", func(t *testing.T) {
		// Arrange
		param = convertToJiraApiResourceParameters(param)
		*param.Username = ""
		context = "ReadIssue"
		issueList = "ABC-123 DEF-456"

		// Act
		param.InitializeAndValidatePostParse(&context, &issueList)

		// Assert
		assert.False(t, param.Meta.AllMandatoryValuesPresent(), "method AllMandatoryValuesPresent() returned true")
		assert.False(t, param.Meta.Ready(), "method Ready() returned true")
	})

	t.Run("app parameters NOT VALID NOR READY from INVALID inputs (MISSING PASSWORD)", func(t *testing.T) {
		// Arrange
		param = convertToJiraApiResourceParameters(param)
		*param.Password = ""
		context = "ReadIssue"
		issueList = "ABC-123 DEF-456"

		// Act
		param.InitializeAndValidatePostParse(&context, &issueList)

		// Assert
		assert.False(t, param.Meta.AllMandatoryValuesPresent(), "method AllMandatoryValuesPresent() returned true")
		assert.False(t, param.Meta.Ready(), "method Ready() returned true")
	})

	t.Run("app parameters NOT VALID NOR READY from INVALID inputs (EMPTY ISSUES)", func(t *testing.T) {
		// Arrange
		param = convertToJiraApiResourceParameters(param)
		param.IssueList = make([]string, 0)
		context = "ReadIssue"
		issueList = ""

		// Act
		param.InitializeAndValidatePostParse(&context, &issueList)

		// Assert
		assert.True(t, param.Meta.AllMandatoryValuesPresent(), "method AllMandatoryValuesPresent() returned false")
		assert.False(t, param.Meta.Ready(), "method Ready() returned true")
	})

}

func convertToJiraApiResourceParameters(param configuration.JiraAPIResourceParameters) configuration.JiraAPIResourceParameters {
	*param.JiraAPIUrl = tParam.JiraAPIUrl
	*param.Username = tParam.Username
	*param.Password = tParam.Password
	*param.Destination = tParam.Destination
	*param.ClosedStatusName = tParam.ClosedStatusName
	*param.TransitionName = tParam.TransitionName
	*param.Destination = tParam.Destination

	*param.EditCustomFieldParam.CustomFieldName = tParam.EditCustomFieldParam.CustomFieldName

	param.Flags.ForceOpen = fPtr

	TestLoggingLevel := "INFO"
	param.LoggingLevel = &TestLoggingLevel

	return param
}
