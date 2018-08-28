// Copyright (c) 2017 Derek Santos. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.
package entigo

import (
	"testing"
)

func TestPassword(t *testing.T) {

	p := Password("test")
	p.Hash()

	l := len(p)
	if l != 60 {
		t.Errorf("length of password should be 60 but was %d", l)
	}

	plain := "test"
	err := p.Compare(plain)
	if err != nil {
		t.Errorf("password comparison failed. %v", err)
	}
}
