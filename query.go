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
)

// QueryRows facilitates queries that retrieve 0 or more rows.
// The scanner fucn argument is expected to handle scanning each row
// as desired, typically into a slice of structs. The args parameter
// is passed into the query for parameterized queries.
func QueryRows(db *sql.DB, query string, scanner func(*sql.Rows) error, args ...interface{}) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = scanner(rows)
		if err != nil {
			return err
		}
	}
	return nil
}

// QueryRow facilitates queries that retrieve a single row.
// The scanner func argument is expected to handle scanning the row
// if it's found. The args parameter is passed into the query for
// parameterized queries.
func QueryRow(db *sql.DB, query string, scanner func(*sql.Row) error, args ...interface{}) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(args...)
	return scanner(row)
}

// Insert executes an insert statement and returns the
// last insert ID or an error.
func Insert(db *sql.DB, query string, args ...interface{}) (int64, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	r, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	return r.LastInsertId()
}

// Exec executes a statemen and returns an
// error if there was one.
func Exec(db *sql.DB, query string, args ...interface{}) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	return err
}
