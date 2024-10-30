package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"table:categories"`
	ID            int `bun:",pk,autoincrement"`
	Name          string
	CreatedAt     time.Time `bun:",nullzero"`
	UpdatedAt     time.Time `bun:",nullzero"`
}
