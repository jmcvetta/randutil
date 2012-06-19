/*
                                    Goutil
                                  Test Suite


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
	"math/rand"
	"testing"
)

// Test that randString produces a string within specified min/max length
// parameters.  The actual randonimity of the string is not tested.
func TestRandAlphanumeric(t *testing.T) {
	min := rand.Intn(100)
	max := min + 1 + rand.Intn(100)
	s := RandAlphanumeric(min, max)
	switch true {
	case len(s) < min:
		t.Error("Random string is too short")
	case len(s) > max:
		t.Error("Random string is too short")
	}
	return
}
