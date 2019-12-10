package configuration

import "testing"

func TestAllMandatoryValuesPresent(t *testing.T) {
	tables := []struct {
		r bool
		p bool
		m bool
	}{
		{true, true, true},
		{false, false, false},
		{false, true, false},
		{false, false, true},
	}

	for _, table := range tables {
		meta := &MetaParameters{parsed: table.p, mandatoryPresent: table.m}
		result := meta.AllMandatoryValuesPresent()
		if result != table.r {
			t.Errorf("Bool value was incorrect, got: %t, want: %t.", result, table.r)
		}
	}
}

func TestReady(t *testing.T) {
	tables := []struct {
		r bool
		p bool
		v bool
	}{
		{true, true, true},
		{false, false, false},
		{false, true, false},
		{false, false, true},
	}

	for _, table := range tables {
		meta := &MetaParameters{parsed: table.p, valid: table.v}
		result := meta.Ready()
		if result != table.r {
			t.Errorf("Bool value was incorrect, got: %t, want: %t.", result, table.r)
		}
	}
}