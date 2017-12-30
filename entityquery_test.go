/*
MIT License

Copyright (c) 2017 Derek Santos

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package entigo

import (
	"encoding/json"
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

	query := EntityCollection(func() EntityDefiner { return &Customer{} })
	test, err := query.Select(db, "WHERE email LIKE '%@test.com'")
	if err != nil {
		t.Error(err)
		return
	}

	if len(test) != 3 {
		t.Error("Number of customers in test.com domain should be 3.")
	}

	b, err := json.Marshal(test)
	if err != nil {
		t.Error(err)
		return
	}

}
