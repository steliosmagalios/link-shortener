// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: links.sql

package repository

import (
	"context"
)

const getAllLinks = `-- name: GetAllLinks :many
SELECT id, location, slug FROM links
`

func (q *Queries) GetAllLinks(ctx context.Context) ([]Link, error) {
	rows, err := q.db.Query(ctx, getAllLinks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(&i.ID, &i.Location, &i.Slug); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
