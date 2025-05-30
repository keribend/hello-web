// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: event.sql

package repository

import (
	"context"
)

const create = `-- name: Create :one
INSERT INTO event (name) VALUES (?)
RETURNING id, name, create_time
`

func (q *Queries) Create(ctx context.Context, name string) (Event, error) {
	row := q.db.QueryRowContext(ctx, create, name)
	var i Event
	err := row.Scan(&i.ID, &i.Name, &i.CreateTime)
	return i, err
}

const findAllEvents = `-- name: FindAllEvents :many
SELECT id, name, create_time FROM event
`

func (q *Queries) FindAllEvents(ctx context.Context) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, findAllEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(&i.ID, &i.Name, &i.CreateTime); err != nil {
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

const findByID = `-- name: FindByID :one
SELECT id, name, create_time FROM event WHERE id = ?
`

func (q *Queries) FindByID(ctx context.Context, id int64) (Event, error) {
	row := q.db.QueryRowContext(ctx, findByID, id)
	var i Event
	err := row.Scan(&i.ID, &i.Name, &i.CreateTime)
	return i, err
}

const findByName = `-- name: FindByName :one
SELECT id, name, create_time FROM event WHERE name = ?
`

func (q *Queries) FindByName(ctx context.Context, name string) (Event, error) {
	row := q.db.QueryRowContext(ctx, findByName, name)
	var i Event
	err := row.Scan(&i.ID, &i.Name, &i.CreateTime)
	return i, err
}

const updateName = `-- name: UpdateName :exec
UPDATE event SET name = ? WHERE id = ?
`

type UpdateNameParams struct {
	Name string
	ID   int64
}

func (q *Queries) UpdateName(ctx context.Context, arg UpdateNameParams) error {
	_, err := q.db.ExecContext(ctx, updateName, arg.Name, arg.ID)
	return err
}
