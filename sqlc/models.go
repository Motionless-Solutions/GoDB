// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"time"
)

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Passwordhash string    `json:"passwordhash"`
	Fullname     string    `json:"fullname"`
	Createdate   time.Time `json:"createdate"`
	Role         int32     `json:"role"`
}
