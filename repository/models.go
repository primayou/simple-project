// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package repository

import (
	"database/sql"
)

type Post struct {
	ID        sql.NullInt64 `json:"id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	Content   string        `json:"content"`
}

type User struct {
	ID           sql.NullInt64 `json:"id"`
	CreatedAt    sql.NullTime  `json:"created_at"`
	UpdatedAt    sql.NullTime  `json:"updated_at"`
	Email        string        `json:"email"`
	PasswordHash string        `json:"password_hash"`
}
