// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: url.sql

package store

import (
	"context"
)

const createURL = `-- name: CreateURL :exec
INSERT INTO "url" (
  public_id, url
) VALUES (
  $1, $2
)
`

type CreateURLParams struct {
	PublicID string
	Url      string
}

func (q *Queries) CreateURL(ctx context.Context, arg CreateURLParams) error {
	_, err := q.db.ExecContext(ctx, createURL, arg.PublicID, arg.Url)
	return err
}

const getCount = `-- name: GetCount :one
SELECT count FROM "url"
WHERE public_id = $1
`

func (q *Queries) GetCount(ctx context.Context, publicID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCount, publicID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getID = `-- name: GetID :one
SELECT public_id FROM "url"
WHERE url = $1 LIMIT 1
`

func (q *Queries) GetID(ctx context.Context, url string) (string, error) {
	row := q.db.QueryRowContext(ctx, getID, url)
	var public_id string
	err := row.Scan(&public_id)
	return public_id, err
}

const getURL = `-- name: GetURL :one
SELECT url, count FROM "url"
WHERE public_id = $1 LIMIT 1
`

type GetURLRow struct {
	Url   string
	Count int64
}

func (q *Queries) GetURL(ctx context.Context, publicID string) (GetURLRow, error) {
	row := q.db.QueryRowContext(ctx, getURL, publicID)
	var i GetURLRow
	err := row.Scan(&i.Url, &i.Count)
	return i, err
}

const updateCount = `-- name: UpdateCount :exec
UPDATE "url"
SET count = $2
WHERE public_id = $1
`

type UpdateCountParams struct {
	PublicID string
	Count    int64
}

func (q *Queries) UpdateCount(ctx context.Context, arg UpdateCountParams) error {
	_, err := q.db.ExecContext(ctx, updateCount, arg.PublicID, arg.Count)
	return err
}
