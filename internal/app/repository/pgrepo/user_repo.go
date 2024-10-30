package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/northwindman/book-shop/internal/app/domain"
	"github.com/northwindman/book-shop/internal/app/repository/models"
	"github.com/northwindman/book-shop/internal/pkg/pg"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r UserRepo) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	dbUser := domainToUser(user)

	var insertedUser models.User
	err := r.db.NewInsert().Model(&dbUser).Returning("*").Scan(ctx, &insertedUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to insert a user: %w", err)
	}

	domainUser, err := userToDomain(insertedUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create domain user: %w", err)
	}

	return domainUser, nil
}

func (r UserRepo) GetUser(ctx context.Context, username string) (domain.User, error) {
	var dbUser models.User
	err := r.db.NewSelect().Model(&dbUser).Where("username = ?", username).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("failed to get a user: %w", err)
	}

	user, err := userToDomain(dbUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create domain user: %w", err)
	}

	return user, nil
}

func (r UserRepo) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	var dbUser models.User
	err := r.db.NewSelect().Model(&dbUser).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("failed to get a user: %w", err)
	}

	user, err := userToDomain(dbUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create domain user: %w", err)
	}

	return user, nil
}
