// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const deleteAlbumByID = `-- name: DeleteAlbumByID :exec
DELETE FROM albums
WHERE id=$1
`

func (q *Queries) DeleteAlbumByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAlbumByID, id)
	return err
}

const getAlbumByID = `-- name: GetAlbumByID :one
SELECT id, title, artist, price 
FROM albums 
WHERE id=$1
`

func (q *Queries) GetAlbumByID(ctx context.Context, id int32) (Album, error) {
	row := q.db.QueryRowContext(ctx, getAlbumByID, id)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Artist,
		&i.Price,
	)
	return i, err
}

const getAlbums = `-- name: GetAlbums :many
SELECT id, title, artist, price FROM albums
`

func (q *Queries) GetAlbums(ctx context.Context) ([]Album, error) {
	rows, err := q.db.QueryContext(ctx, getAlbums)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Album
	for rows.Next() {
		var i Album
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Artist,
			&i.Price,
		); err != nil {
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

const postAlbums = `-- name: PostAlbums :one
INSERT INTO albums (title, artist, price)
VALUES ($1, $2, $3)
RETURNING id
`

type PostAlbumsParams struct {
	Title  sql.NullString
	Artist sql.NullString
	Price  sql.NullInt32
}

func (q *Queries) PostAlbums(ctx context.Context, arg PostAlbumsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, postAlbums, arg.Title, arg.Artist, arg.Price)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateAlbumByID = `-- name: UpdateAlbumByID :one
UPDATE albums
SET title = $2, artist = $3, price = $4 
WHERE id=$1
RETURNING id, title, artist, price
`

type UpdateAlbumByIDParams struct {
	ID     int32
	Title  sql.NullString
	Artist sql.NullString
	Price  sql.NullInt32
}

func (q *Queries) UpdateAlbumByID(ctx context.Context, arg UpdateAlbumByIDParams) (Album, error) {
	row := q.db.QueryRowContext(ctx, updateAlbumByID,
		arg.ID,
		arg.Title,
		arg.Artist,
		arg.Price,
	)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Artist,
		&i.Price,
	)
	return i, err
}