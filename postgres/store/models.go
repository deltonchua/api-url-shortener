// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package store

import (
	"time"
)

type Url struct {
	ID        int64
	PublicID  string
	Url       string
	Count     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
