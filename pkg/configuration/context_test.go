package configuration

import "testing"

func TestStringUnknown1(t *testing.T) {
	val := Context(-1).String()
	if val != "Unknown1" {
		t.Errorf("String value was incorrect, got: %s, want: %s.", val, "Unknown")
	}
}
