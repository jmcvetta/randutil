// Copyright (c) 2012 Jason McVetta.  This is Free Software, released under the 
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package randutil

import (
	"fmt"
	"math/rand"
	"testing"
)

// Test that AlphaStringRange produces a string within specified min/max length
// parameters.  The actual randonimity of the string is not tested.
func TestAlphaStringRange(t *testing.T) {
	min := rand.Intn(100)
	max := min + 1 + rand.Intn(100)
	s, err := AlphaStringRange(min, max)
	if err != nil {
		t.Error(err)
	}
	switch true {
	case len(s) < min:
		t.Error("Random string is too short")
	case len(s) > max:
		t.Error("Random string is too short")
	}
	return
}

// Test that IntRange produces an integer between min and max
func TestIntRange(t *testing.T) {
	min := 567
	max := 890
	i, err := IntRange(min, max)
	if err != nil {
		t.Error(err)
	}
	if i > max || i < min {
		t.Error("IntRange returned an out-of-range integer")
	}
	// Check that we get an error when min > max
	i, err = IntRange(max, min)
	if err != MinMaxError {
		msg := fmt.Sprintf("Expected error when min > max, but got:", err)
		t.Error(msg)
	}
}

// Test that the strings we produce are actually random.  This is done by 
// comparing two 50,000 character generated random strings and checking that
// they differ.  It is quite unlikely, but not strictly impossible, that two
// truly random strings will be identical.
func TestRandonimity(t *testing.T) {
	l := 50000
	s1, err := AlphaString(l)
	if err != nil {
		t.Error(err)
	}
	s2, err := AlphaString(l)
	if err != nil {
		t.Error(err)
	}
	if s1 == s2 {
		msg := "Generated two identical 'random' strings - this is probably an error"
		t.Error(msg)
	}
}
