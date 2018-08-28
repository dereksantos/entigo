// Copyright (c) 2017 Derek Santos. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.
package entigo

import "testing"

func TestSha256(t *testing.T) {
	actual := Sha256("test")
	actual.Hash()

	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	if string(actual) != expected {
		t.Errorf("Sha256 value expected %s but got %s", expected, actual)
	}
}
