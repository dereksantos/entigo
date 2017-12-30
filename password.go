// Copyright (c) 2017 Derek Santos. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.
package entigo

import (
	"database/sql/driver"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Password is a string type used to represent
// secure password fields that are hashed using BCrypt.
type Password string

// Hash will bcrypt the pointer receiver's value and
// set it through the pointer receiver. It will return
// an error if bcrypt fails.
func (p *Password) Hash() error {
	str := string(*p)
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err == nil {
		*p = Password(hashed)
	}
	return err
}

// Compare will use bcrypt to check for equality between
// the specified candidate parameter and the receiver.
// It will return an error if they do not match.
func (p *Password) Compare(candidate string) error {
	return bcrypt.CompareHashAndPassword([]byte(*p), []byte(candidate))
}

func (p *Password) Scan(value interface{}) error {
	v, ok := value.(string)
	if ok {
		*p = Password(v)
		return nil
	}
	return fmt.Errorf("Can't convert %T to *entity.Password", value)
}

func (p Password) Value() (driver.Value, error) {
	return string(p), nil
}
