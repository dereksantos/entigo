// Copyright (c) 2017 Derek Santos. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.
package entigo

import (
	"database/sql"
	"fmt"
	"strings"
)

type EntityCollection func() EntityDefiner

func (c EntityCollection) Select(db *sql.DB, clause string, args ...interface{}) ([]EntityDefiner, error) {
	e := c().Entity()
	fields := e.Fields
	names := []string{e.Key.Name}
	for _, f := range fields {
		names = append(names, f.Name)
	}

	query := fmt.Sprintf("SELECT %s FROM %s %s", strings.Join(names, ","), e.Name, clause)

	entities := []EntityDefiner{}
	err := QueryRows(db, query, func(rows *sql.Rows) error {
		def := c()
		e := def.Entity()
		fields := e.Fields
		values := []interface{}{e.Key.Value}
		for _, f := range fields {
			values = append(values, f.Value)
		}

		err := rows.Scan(values...)
		if err != nil {
			return err
		}
		entities = append(entities, def)
		return nil
	}, args...)
	return entities, err
}
