package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Book is a domain book.
type Book struct {
	bun.BaseModel `bun:"table:books"`
	ID            int `bun:",pk,autoincrement"`
	Title         string
	Year          int
	Author        string
	Price         int
	Stock         int
	CategoryID    int
	CreatedAt     time.Time `bun:",nullzero"`
	UpdatedAt     time.Time `bun:",nullzero"`
}
