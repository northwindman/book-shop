package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Cart struct {
	bun.BaseModel `bun:"table:carts"`
	UserID        int       `bun:"user_id"`
	BookIDs       []int     `bun:"book_ids,array"`
	CreatedAt     time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero"`
}
