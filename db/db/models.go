// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
)

type Album struct {
	ID     int32
	Title  sql.NullString
	Artist sql.NullString
	Price  sql.NullInt32
}
