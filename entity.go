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
	"database/sql"
	"fmt"
	"strings"
)

// An Entity defines the relationship between an instance of a
// struct and a row in a database table or collection.
// It's important to note that it's related to an instance, and so
// has pointers to the stucts fields. This facilitates row scanning
// and prepared statements.
type Entity struct {
	Name   string
	Key    *Field
	Fields []*Field
}

// A Field maps the name of a database column to a pointer in a struct.
// If ReadOnly is true, this field will not be included in writes to
// the database.
type Field struct {
	Name     string
	Value    interface{}
	ReadOnly bool

	// NonIncrementing should be set to true when you need a primary key
	// whos value does not auto increment.
	NonIncrementing bool
}

// An EntityDefiner must return a pointer to an Entity struct. This
// interface should be implemented by structs that relate to a database
// table or collection.
type EntityDefiner interface {
	Entity() *Entity
}

// Get scans a single row of data from an auto-generated sql statement
// into the Entity's value pointers. It returns the row based on the
// primary key value of the entity.
//
// Example:
// 		person := &Person{ID:1}
// 		err := person.Entity().Get()
//
func (ent *Entity) Get(db *sql.DB) error {
	q := "SELECT %s FROM %s WHERE %s=?"
	fields := ent.Fields
	names := make([]string, len(fields))
	values := make([]interface{}, len(fields))
	for i, f := range fields {
		names[i] = f.Name
		values[i] = f.Value
	}
	q = fmt.Sprintf(q, strings.Join(names, ","), ent.Name, ent.Key.Name)
	return QueryRow(db, q, func(row *sql.Row) error {
		return row.Scan(values...)
	}, ent.Key.Value)
}

// Insert executes an insert statement on a single row of data using the
// values in the receiver. The new ID will be returned if applicable.
// Fields with ReadOnly set to true will not be included in the insert.
//
// Example:
// 		person := &Person{Name:"John Doe", Email"test@test.com"}
// 		id, err := person.Entity().Insert()
// 		if err != nil {
// 			panic(err)
// 		}
// 		person.ID = id
//
func (ent *Entity) Insert(db *sql.DB) (int64, error) {
	q := "INSERT INTO %s(%s) VALUES (%s)"
	fields := ent.Fields
	names := []string{}
	values := []interface{}{}

	if ent.Key.NonIncrementing {
		names = append(names, ent.Key.Name)
		values = append(values, ent.Key.Value)
	}

	for _, f := range fields {
		if !f.ReadOnly {
			names = append(names, f.Name)
			values = append(values, f.Value)
		}
	}

	columns := strings.Join(names, ",")
	params := strings.TrimSuffix(strings.Repeat("?,", len(names)), ",")
	s := fmt.Sprintf(q, ent.Name, columns, params)
	return Insert(db, s, values...)
}

// Update executes an update statement on  a single row of data using
// values in the receiver. Fields with ReadOnly set to true will not be
// included in the update.
//
// Example:
//
// 		person := &Person{ID:1}
// 		e := person.Entity()
// 		err := e.Get()
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		person.Name = "New Name"
// 		err = e.Update()
// 		if err != nil {
// 			panic(err)
// 		}
//
func (ent *Entity) Update(db *sql.DB) error {
	fields := ent.Fields
	sets := []string{}
	values := []interface{}{}

	if ent.Key.NonIncrementing {
		sets = append(sets, fmt.Sprintf("%s=?", ent.Key.Name))
		values = append(values, ent.Key.Value)
	}

	for _, f := range fields {
		if !f.ReadOnly {
			sets = append(sets, fmt.Sprintf("%s=?", f.Name))
			values = append(values, f.Value)
		}
	}

	values = append(values, ent.Key.Value)

	q := "UPDATE %s SET %s WHERE %s = ?"
	q = fmt.Sprintf(q, ent.Name, strings.Join(sets, ","), ent.Key.Name)
	return Exec(db, q, values...)
}

// Delete executes a delete statement using the Key value of
// the receiver. The entity requires only the Key to be set
// for the delete to work.
//
// Example:
//
//		person := &Person{ID:1}
//		err := person.Entity().Delete()
//		if err != nil {
//			panic(err)
//		}
//
func (ent *Entity) Delete(db *sql.DB) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE %s=?", ent.Name, ent.Key.Name)
	return Exec(db, q, ent.Key.Value)
}
