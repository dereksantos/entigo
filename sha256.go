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
	"crypto/sha256"
	"database/sql/driver"
	"fmt"
)

type Sha256 string

func (g *Sha256) Hash() {
	data := []byte(*g)
	hash := sha256.Sum256(data)
	val := Sha256(fmt.Sprintf("%x", hash))
	*g = val
}

func (p *Sha256) Scan(value interface{}) error {
	v, ok := value.(string)
	if ok {
		*p = Sha256(v)
		return nil
	}
	return fmt.Errorf("Can't convert %T to *entity.Sha256", value)
}

func (p Sha256) Value() (driver.Value, error) {
	return string(p), nil
}
