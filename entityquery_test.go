// Copyright (c) 2017 Derek Santos. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.
package entigo

import (
	"testing"
	"time"
)

func TestCollectionSelect(t *testing.T) {

	inserts := []*Customer{
		&Customer{
			Name:    "John",
			Email:   "john+1@test.com",
			Created: time.Now(),
			Updated: time.Now(),
		},
		&Customer{
			Name:    "John2",
			Email:   "john+2@test.com",
			Created: time.Now(),
			Updated: time.Now(),
		},
		&Customer{
			Name:    "John3",
			Email:   "john+3@test.com",
			Created: time.Now(),
			Updated: time.Now(),
		},
		&Customer{
			Name:    "John4",
			Email:   "john+4@other.com",
			Created: time.Now(),
			Updated: time.Now(),
		},
	}
	for _, v := range inserts {
		_, err := v.Entity().Insert(db)
		if err != nil {
			t.Error(err)
			return
		}
	}

	collection := EntityCollection(func() EntityDefiner { return &Customer{} })

	test, err := collection.Select(db, "WHERE email LIKE '%@test.com'")
	if err != nil {
		t.Error(err)
		return
	}

	if len(test) != 3 {
		t.Error("Number of customers in test.com domain should be 3.")
	}

	test, err = collection.Select(db, "WHERE email = ?", "john+4@other.com")
	if err != nil {
		t.Error(err)
		return
	}

	if len(test) != 1 {
		t.Error("Should have one row")
	}

}
