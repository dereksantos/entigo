// Copyright (c) 2017 Derek Santos. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.
package entigo

import "testing"

func TestMD5(t *testing.T) {
	actual := MD5(`Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`)
	actual.Hash()

	ptr := new(MD5)
	*ptr = MD5("test")
	ptr.Hash()

	t.Logf("%s", *ptr)

	expected := "bdf55c952333a4f0992f429e152f03f7"
	if string(actual) != expected {
		t.Errorf("MD5 value expected %s but got %s", expected, actual)
	}
}
