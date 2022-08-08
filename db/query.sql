-- name: GetAlbums :many
SELECT id, title, artist, price FROM albums;

-- name: GetAlbumByID :one
SELECT id, title, artist, price 
FROM albums 
WHERE id=$1;

-- name: DeleteAlbumByID :exec
DELETE FROM albums
WHERE id=$1;

-- name: PostAlbums :one
INSERT INTO albums (title, artist, price)
VALUES ($1, $2, $3)
RETURNING id;

-- name: UpdateAlbumByID :one
UPDATE albums
SET title = $2, artist = $3, price = $4 
WHERE id=$1
RETURNING *;