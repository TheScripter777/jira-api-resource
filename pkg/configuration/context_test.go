package configuration

import "testing"

func TestString(t *testing.T) {
	tables := []struct {
		i int
		o string
	}{
		{-10, "Unknown"},
		{-1, "Unknown"},
		{0, "ReadIssue"},
		{1, "EditCustomField"},
		{999, "Unknown"},
	}

	for _, table := range tables {
		val := Context(table.i).String()
		if val != table.o {
			t.Errorf("String value was incorrect, got: %s, want: %s.", val, table.o)
		}
	}
}

func TestGetContext(t *testing.T) {
	var tables = []struct {
		i int
		o string
	}{
		{0, "ReadIssue"},
		{1, "EditCustomField"},
		{2, "Unknown"},
	}

	for _, table := range tables {
		ctx := GetContext(table.o)
		if ctx != Context(table.i) {
			t.Errorf("Context (int) value was incorrect, got: %d, want: %d.", ctx, table.i)
		}
	}
}
