// See parameters.go for this package's comment
package configuration

// Context is a simple integer to facilitate the handling of various context names via an enum-like strategy.
type Context int

/*
This is an enum-like constant block to define every available context of execution.
TODO: update comments
*/
const (
	ReadIssue Context = iota
	EditCustomField
	Unknown
)

var names = [...]string{"ReadIssue", "EditCustomField", "Unknown"}

// Returns the string value of the current Context
func (c Context) String() string {

	if c < ReadIssue || c >= Unknown {
		return "Unknown"
	}

	return names[c]

}

// This function is the implementation of the 'valueOf' mechanic of an enum-like construct
func GetContext(contextString string) Context {
	for index, name := range names {
		if contextString == "Unknown" {
			return Context(len(names) - 1)
		}

		if name == contextString {
			return Context(index)
		}
	}

	return Unknown
}
