/*
                                    Goutil
             A Motley Assortment of Utilities for the Go Language


@author: Jason McVetta <jason.mcvetta@gmail.com>
@copyright: (c) 2012 Jason McVetta
@license: GPL v3 - http://www.gnu.org/copyleft/gpl.html

********************************************************************************
Goutil is free software: you can redistribute it and/or modify it under the
terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.

Goutil is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
Goutil.  If not, see <http://www.gnu.org/licenses/>.
********************************************************************************

*/

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
