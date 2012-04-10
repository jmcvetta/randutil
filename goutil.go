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

package goutil

import (
	"crypto/rand"
	"math/big"
)

const (
	// Set of characters to use for generating random strings
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 abcdefghijklmnopqrstuvwxyz" + 
	"~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

// randString returns a random string no more than at least min and no more 
// than max characters long.
func randString(min, max int) string {
	//
	// First determine the length of string to be generated
	//
	var err error  // Holds errors
	var b *big.Int // Holds random bigints
	var r int      // Holds random integers
	maxRand := max - min
	b, err = rand.Int(rand.Reader, big.NewInt(int64(maxRand)))
	if err != nil {
		panic(err) // WTF?
	}
	r = int(b.Int64())
	strlen := min + r
	charlen := len(chars)
	randstr := ""
	for i := 0; i < strlen; i++ {
		b, err = rand.Int(rand.Reader, big.NewInt(int64(charlen-1)))
		if err != nil {
			panic(err) // WTF?
		}
		r = int(b.Int64())
		c := string(chars[r])
		randstr += c
	}
	return randstr
}

