// Copyright (c) 2017 Derek Santos. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.
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
