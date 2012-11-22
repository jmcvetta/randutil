// Copyright (c) 2012 Jason McVetta.  This is Free Software, released under the 
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

// Package randutil provides various convenience functions for dealing with 
// random numbers and strings.
package randutil

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	// Set of characters to use for generating random strings
	Alphanumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz"
	Ascii        = Alphanumeric + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

var MinMaxError = errors.New("Min cannot be greater than max.")

// IntRange returns a random integer in the range from min to max.
func IntRange(min, max int) (int, error) {
	var result int
	switch {
	case max == min:
		result = max
	case max > min:
		// Choose a random string lenth between min and max
		maxRand := max - min
		b, err := rand.Int(rand.Reader, big.NewInt(int64(maxRand)))
		if err != nil {
			return result, err
		}
		result = min + int(b.Int64())
	case min > max:
		// Fail with error
		return result, MinMaxError
	}
	return result, nil
}

// String returns a random string n characters long, composed of entities 
// from charset.
func String(n int, charset string) (string, error) {
	var randstr string // Random string to return
	charlen := len(charset)
	// This is probably not the most efficient algorithm
	for i := 0; i < n; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(int64(charlen-1)))
		if err != nil {
			return randstr, err
		}
		r := int(b.Int64())
		c := string(charset[r])
		randstr += c
	}
	return randstr, nil
}

// StringRange returns a random string at least min and no more than max 
// characters long, composed of entitites from charset.
func StringRange(min, max int, charset string) (string, error) {
	//
	// First determine the length of string to be generated
	//
	var err error      // Holds errors
	var strlen int     // Length of random string to generate
	var randstr string // Random string to return
	strlen, err = IntRange(min, max)
	if err != nil {
		return randstr, err
	}
	randstr, err = String(strlen, charset)
	if err != nil {
		return randstr, err
	}
	return randstr, nil
}

// AlphaRange returns a random alphanumeric string at least min and no more 
// than max characters long.
func AlphaStringRange(min, max int) (string, error) {
	return StringRange(min, max, Alphanumeric)
}

// AlphaString returns a random alphanumeric string n characters long.
func AlphaString(n int) (string, error) {
	return String(n, Alphanumeric)
}

// ChoiceString returns a random selection from an array of strings.
func ChoiceString(choices []string) (string, error) {
	var winner string
	length := len(choices)
	i, err := IntRange(0, length)
	winner = choices[i]
	return winner, err
}

// ChoiceInt returns a random selection from an array of integers.
func ChoiceInt(choices []int) (int, error) {
	var winner int
	length := len(choices)
	i, err := IntRange(0, length)
	winner = choices[i]
	return winner, err
}

type Choice struct {
	Weight int
	Item   interface{}
}

func WeightedChoice(choices []Choice) (Choice, error) {
	// Based on this algorithm:
	//     http://eli.thegreenplace.net/2010/01/22/weighted-random-generation-in-python/
	var ret Choice
	sum := 0
	for _, c := range choices {
		sum += c.Weight
	}
	r, err := IntRange(0, sum)
	if err != nil {
		return ret, err
	}
	for _, c := range choices {
		r -= c.Weight
		if r < 0 {
			return c, nil
		}
	}
	err = errors.New("Internal error - code should not reach this point")
	return ret, err
}
