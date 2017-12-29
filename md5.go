package entigo

import (
	"crypto/md5"
	"database/sql/driver"
	"fmt"
)

type MD5 string

func (g *MD5) Hash() {
	data := []byte(*g)
	hash := md5.Sum(data)
	val := MD5(fmt.Sprintf("%x", hash))
	*g = val
}

func (p *MD5) Scan(value interface{}) error {
	v, ok := value.(string)
	if ok {
		*p = MD5(v)
		return nil
	}
	return fmt.Errorf("Can't convert %T to *entity.MD5", value)
}

func (p MD5) Value() (driver.Value, error) {
	return string(p), nil
}
