// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const getAll = `-- name: GetAll :many
SELECT first_name, last_name, age FROM users
`

func (q *Queries) GetAll(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.FirstName, &i.LastName, &i.Age); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
