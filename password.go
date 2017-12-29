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
