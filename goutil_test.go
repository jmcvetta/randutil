package goutil

import (
    "testing"
    "math/rand"
)

// Test that randString produces a string within specified min/max length
// parameters.  The actual randonimity of the string is not tested.
func TestRandString(t *testing.T) {
	min := rand.Intn(100)
	max := min + 1 + rand.Intn(100)
	s := RandString(min, max)
	switch true {
	case len(s) < min:
		t.Error("Random string is too short")
	case len(s) > max:
		t.Error("Random string is too short")
	}
	return
}

