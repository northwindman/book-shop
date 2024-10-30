package models

import (
	"time"

	"github.com/uptrace/bun"
)

// User is a domain user.
type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int `bun:",pk,autoincrement"`
	Username      string
	Password      string
	Admin         bool
	CreatedAt     time.Time `bun:",nullzero"`
	UpdatedAt     time.Time `bun:",nullzero"`
}
