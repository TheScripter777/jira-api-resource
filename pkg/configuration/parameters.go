/*
Package configuration provides the structs that contains this application's Go flags
and other custom parameters determined from the Go flags. Parsing of said flags is
done via the JiraAPIResourceParameters.Parse() method

To initialize the JiraAPIResourceParameters one only needs to do the following:

	// Short declaration of the variable
	params := configuration.JiraAPIResourceParameters{}

	// Execution of the 'Parse' method
	params.Parse()

After that specific method, the flags.Parse method of the Go lang flags api will have been called.

The package also defines a wide array of constants representing the flags name in the command line.
*/
package configuration

import (
	"flag"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/auth"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/helpers"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/log"
	"strings"
)

// Definition of constants that are use for the flags setup
const (
	// Parameters
	jiraAPIURL               = "url"
	username                 = "username"
	password                 = "password"
	destination              = "destination"
	context                  = "context"
	issueList                = "issues"
	customFieldName          = "customFieldName"
	customFieldType          = "customFieldType"
	customFieldValueAsIs     = "customFieldValue"
	customFieldValueFromFile = "customFieldValueFromFile"
	loggingLevel             = "loggingLevel"
	closedStatusName         = "closedStatusName"
	transitionName           = "transitionName"
	commentBody              = "commentBody"

	// Flags
	forceOnParent    = "forceOnParent"
	forceOpen        = "forceOpen"
	keepGoingOnError = "keepGoing"

	// Default values and descriptions for both paramaters and flags
	jiraAPIURLDefault                   = ""
	jiraAPIURLDescription               = "The base URL of the Jira API"
	usernameDefault                     = ""
	usernameDescription                 = "The username used to connect to the Jira API"
	passwordDefault                     = ""
	passwordDescription                 = "The password needed to connect to the Jira API"
	destinationDefault                  = "./"
	destinationDescription              = "The destination to output new version(s) when using the 'in'"
	contextDefault                      = ""
	contextDescription                  = "The context of execution. {'ReadIssue', 'EditCustomField'}"
	issueListDefault                    = ""
	issueListDescription                = "The issue or list of issues to execute the specified context to"
	customFieldNameDefault              = ""
	customFieldNameDescription          = "Certain operation (such as edits) might require the user to specify the name of the custome field so that the resource may find the appropriate custom field"
	customFieldTypeDefault              = "string"
	customFieldTypeDescription          = "The type that is required by the field via the Jira API"
	customFieldValueAsIsDefault         = ""
	customFieldValueAsIsDescription     = "The value of the field that will be updated (in case of update workflow)"
	customFieldValueFromFileDefault     = ""
	customFieldValueFromFileDescription = "The value of the field, stored in a file, that will be updated (in case of update workflow)"
	loggingLevelDefault                 = "INFO"
	loggingLevelDescription             = "The level of the loggers of the application {'ALL', 'DEBUG', 'ERROR', 'INFO', 'WARN', 'OFF'}"
	closedStatusNameDefault             = "Closed"
	closedStatusNameDescription         = "The written value of the closed status as displayed in the Jira issue. Used for openness check."
	transitionNameDefault               = "Reopened"
	transitionNameDescription           = "The name (as written in Jira) of the desired nwe status."
	commentBodyDefault                  = ""
	commentBodyDescription              = "The text body of the comment that will be posted to specified issue(s)."
	_                                   = /*forceOnParentDefault*/ false
	forceOnParentDescription            = "Flag that indicates if we want to force all operation on the parent issue (if there's one)"
	_                                   = /*forceOpenDefault*/ false
	forceOpenDescription                = "Flag that will, if the issue if closed, forcefully open it, apply the changes and then close it back"
	_                                   = /*keepGoingOnErrorDefault*/ false
	keepGoingOnErrorDescription         = "Flag that makes the pipeline continue even if an error occurs on a specific issue."
)

// JiraAPIResourceParameters is a struct that holds every possible parameters/flags known by the application.
// They are all parsed via the flag.Parse() method of the Go flags api. The struct also contains the meta-parameters
// (see meta-parametes.go).
type JiraAPIResourceParameters struct {
	JiraAPIUrl       *string
	Username         *string
	Password         *string
	Destination      *string
	Context          Context
	IssueList        []string
	LoggingLevel     *string
	ClosedStatusName *string
	TransitionName   *string

	ReadIssueParam       JiraApiResourceParametersReadIssue
	EditCustomFieldParam JiraApiResourceParametersEditCustomField
	AddComment           JiraApiResourceParametersAddComment

	ActiveIssue string         // The **SINGLE** issue that the resource is currently processing
	Meta        MetaParameters //
	Flags       JiraAPIResourceFlags
}

// This struct is used to separate the parameters (which takes values in the command line) of the flags (which don't).
// That being said the values in this struct are still parsed via the Go flags api. They've been put 'aside' for clariry
// purposes.
type JiraAPIResourceFlags struct {
	ForceOnParent    *bool
	ForceOpen        *bool
	KeepGoingOnError *bool
}

type JiraApiResourceParametersReadIssue struct {
}

type JiraApiResourceParametersReadStatus struct {
}

type JiraApiResourceParametersEditCustomField struct {
	CustomFieldName          *string
	CustomFieldType          *string
	CustomFieldValue         *string
	CustomFieldValueFromFile *string
}

type JiraApiResourceParametersAddComment struct {
	CommentBody *string
}

// Method that initialize every parameters/flags and makes the actual call the flag.Parse().
func (param *JiraAPIResourceParameters) Parse() (*string, *string) {
	var contextString *string
	var issueListString *string

	// Parsing parameters
	param.JiraAPIUrl = flag.String(jiraAPIURL, jiraAPIURLDefault, jiraAPIURLDescription)
	param.Username = flag.String(username, usernameDefault, usernameDescription)
	param.Password = flag.String(password, passwordDefault, passwordDescription)
	param.Destination = flag.String(destination, destinationDefault, destinationDescription)
	contextString = flag.String(context, contextDefault, contextDescription)
	issueListString = flag.String(issueList, issueListDefault, issueListDescription)
	param.EditCustomFieldParam.CustomFieldName = flag.String(customFieldName, customFieldNameDefault, customFieldNameDescription)
	param.EditCustomFieldParam.CustomFieldType = flag.String(customFieldType, customFieldTypeDefault, customFieldTypeDescription)
	param.EditCustomFieldParam.CustomFieldValue = flag.String(customFieldValueAsIs, customFieldValueAsIsDefault, customFieldValueAsIsDescription)
	param.EditCustomFieldParam.CustomFieldValueFromFile = flag.String(customFieldValueFromFile, customFieldValueFromFileDefault, customFieldValueFromFileDescription)
	param.AddComment.CommentBody = flag.String(commentBody, commentBodyDefault, commentBodyDescription)

	param.LoggingLevel = flag.String(loggingLevel, loggingLevelDefault, loggingLevelDescription)
	param.ClosedStatusName = flag.String(closedStatusName, closedStatusNameDefault, closedStatusNameDescription)
	param.TransitionName = flag.String(transitionName, transitionNameDefault, transitionNameDescription)

	// Parsing flags
	param.Flags.ForceOnParent = flag.Bool(forceOnParent, false, forceOnParentDescription)
	param.Flags.ForceOpen = flag.Bool(forceOpen, false, forceOpenDescription)
	param.Flags.KeepGoingOnError = flag.Bool(keepGoingOnError, false, keepGoingOnErrorDescription)

	if !param.Meta.parsed {
		flag.Parse()
		param.Meta.parsed = flag.Parsed()
	}

	return contextString, issueListString
}

// TODO: ... A few custom operation are performed afterward such as initialization and validation.
func (param *JiraAPIResourceParameters) InitializeAndValidatePostParse(context, issueList *string) {
	param.initializeContext(context)
	param.initializeIssueList(issueList)
	param.initLogger()
	param.validate()
}

func (param *JiraAPIResourceParameters) validate() {
	// By default both meta flags are set to true
	param.Meta.mandatoryPresent = true
	param.Meta.valid = true

	if *param.JiraAPIUrl == "" || *param.Username == "" || *param.Password == "" || helpers.IsStringPtrNilOrEmtpy(param.Destination) {
		// In this case we are missing one or more mandatory parameters
		// This also causes the input parameters to not be valid
		param.Meta.mandatoryPresent = false
		param.Meta.valid = false
	} else if param.IssueList == nil || len(param.IssueList) == 0 {
		// This case is either
		//   - A nil issue list
		//   - An issue list that was passed but is empty
		//   - A nil or empty project key
		// All of those need to cause an invalid state
		param.Meta.valid = false
	} else if helpers.IsBoolPtrTrue(param.Flags.ForceOpen) && helpers.IsStringPtrNilOrEmtpy(param.TransitionName) {
		param.Meta.valid = false
		param.Meta.Msg = fmt.Sprintf("'%s' flag was specified yet the '%s' parameter was not. This is an invalid configuration.", forceOpen, transitionName)
	}

	auth.Username = *param.Username
	auth.Password = *param.Password

	if param.Meta.mandatoryPresent && param.Meta.valid {
		// Next the initialized context needs to be validated against the input parameters
		// At this point we know we have a valid:
		//	- Context
		//	- Set of stand-alone input parameters
		switch param.Context {
		case Unknown:
			// The specified context wasn't recognized, therefore it isn't valid
			param.Meta.valid = false
			param.Meta.Msg = "Unknown context"
		case EditCustomField:
			if helpers.IsStringPtrNilOrEmtpy(param.EditCustomFieldParam.CustomFieldValue) && helpers.IsStringPtrNilOrEmtpy(param.EditCustomFieldParam.CustomFieldValueFromFile) {
				// In the context of editing a custom field, there's an absolute need to have the value passed as input
				// Here we detected that nothing was passed. The meta-parameter 'valid' must then be set to false
				param.Meta.valid = false
				param.Meta.Msg = fmt.Sprintf("Missing 'custom field' parameter (%s or %s)",
					customFieldValueAsIs,
					customFieldValueFromFile)
			}
		case ReadStatus:
			if helpers.IsStringPtrNilOrEmtpy(param.Destination) {
				// This context requires a destination to store the ouput on concourse
				param.Meta.valid = false
				param.Meta.Msg = fmt.Sprintf("Missing destination")
			}
		case AddComment:
			fallthrough
		case ReadIssue:
			fallthrough
		default:
			param.Meta.valid = true
		}
	}
}

func (param *JiraAPIResourceParameters) initializeContext(contextString *string) {
	if contextString == nil {
		param.Context = Unknown
	} else {
		param.Context = GetContext(*contextString)
	}
}

func (param *JiraAPIResourceParameters) initializeIssueList(issueListString *string) {
	if *issueListString != "" {
		// Clean list to make sure we don't have [,],{,},(,) characters
		issueListStringCleaned := helpers.CleanString(*issueListString)

		param.IssueList = strings.Fields(issueListStringCleaned)

		param.IssueList = helpers.CleanStringSlice(param.IssueList)

		// More than 1 issue specified will set the 'Multiple' flag to true
		param.Meta.MultipleIssue = len(param.IssueList) > 1
	}
}

func (param *JiraAPIResourceParameters) initLogger() {
	log.Logger = log.ResourceLogger{}
	log.Logger.InitLoggerFromParam(*param.LoggingLevel)
}
