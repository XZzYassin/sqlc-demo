// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package repo

import (
	"database/sql"
)

type User struct {
	ID        int32
	Name      string
	Email     string
	Status    bool
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
