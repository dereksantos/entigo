// Copyright (c) 2017 Derek Santos. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.
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
