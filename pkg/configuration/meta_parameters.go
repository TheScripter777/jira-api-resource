// See parameters.go for this package's comment
package configuration

// Meta parameters are parameters not directly specified by the user's input. They indicate the completeness of
// the input parameters and the readiness of the application to begin its execution.
type MetaParameters struct {
	parsed           bool // Indicates whether the flag.Parse() method was called or not
	mandatoryPresent bool // Indicates whether the mandatory parameters are persent or not
	valid            bool // Indicates if the values received for each flags are valid or not

	Msg           string // TODO
	MultipleIssue bool   // Indicates whether the resource received multiple Jira issue to process
}

// Returns true if the flags were parsed by the Go flags api (meta.parsed) and if all the mandatory parameters
// have been assigned a value (meta.mandatoryPresent).
func (meta *MetaParameters) AllMandatoryValuesPresent() bool {
	return meta.parsed && meta.mandatoryPresent
}

// Returns true if the flags were parsed by the Go flags api (meta.parsed) and if the set of parameters received
// forms a valid set. Validation is done via the validate method in the parameters.go file of the configuration package.
func (meta *MetaParameters) Ready() bool {
	return meta.parsed && meta.valid
}
