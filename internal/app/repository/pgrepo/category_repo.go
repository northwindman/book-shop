package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/northwindman/book-shop/internal/app/domain"
	"github.com/northwindman/book-shop/internal/app/repository/models"
	"github.com/northwindman/book-shop/internal/pkg/pg"
)

type CategoryRepo struct {
	db *pg.DB
}

func NewCategoryRepo(db *pg.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r CategoryRepo) GetCategory(ctx context.Context, id int) (domain.Category, error) {
	if id == 0 {
		return domain.Category{}, fmt.Errorf("%w: id", domain.ErrRequired)
	}

	var category models.Category
	err := r.db.NewSelect().Model(&category).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Category{}, domain.ErrNotFound
		}
		return domain.Category{}, fmt.Errorf("failed to get a category: %w", err)
	}

	domainCategory, err := categoryToDomain(category)
	if err != nil {
		return domain.Category{}, fmt.Errorf("failed to create domain category: %w", err)
	}

	return domainCategory, nil
}

func (r CategoryRepo) CreateCategory(ctx context.Context, category domain.Category) (domain.Category, error) {
	dbCategory := domainToCategory(category)

	var insertedCategory models.Category
	err := r.db.NewInsert().Model(&dbCategory).Returning("*").Scan(ctx, &insertedCategory)
	if err != nil {
		return domain.Category{}, fmt.Errorf("failed to insert a category: %w", err)
	}

	domainCategory, err := categoryToDomain(insertedCategory)
	if err != nil {
		return domain.Category{}, fmt.Errorf("failed to create domain category: %w", err)
	}

	return domainCategory, nil

}

func (r CategoryRepo) UpdateCategory(ctx context.Context, category domain.Category) (domain.Category, error) {
	dbCategory := domainToCategory(category)
	dbCategory.UpdatedAt = time.Now()

	var updatedCategory models.Category
	err := r.db.NewUpdate().
		Model(&dbCategory).
		Where("id = ?", dbCategory.ID).
		ExcludeColumn("created_at").
		Returning("*").
		Scan(ctx, &updatedCategory)
	if err != nil {
		return domain.Category{}, fmt.Errorf("failed to update a category: %w", err)
	}

	domainCategory, err := categoryToDomain(updatedCategory)
	if err != nil {
		return domain.Category{}, fmt.Errorf("failed to create domain category: %w", err)
	}

	return domainCategory, nil
}

func (r CategoryRepo) DeleteCategory(ctx context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf("%w: id", domain.ErrRequired)
	}

	_, err := r.db.NewDelete().Model((*models.Category)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete a category: %w", err)
	}

	return nil
}

func (r CategoryRepo) GetCategories(ctx context.Context) ([]domain.Category, error) {
	var categories []models.Category
	err := r.db.NewSelect().Model(&categories).Order("id").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to select categories: %w", err)
	}

	var domainCategories []domain.Category
	for _, category := range categories {
		domainCategory, err := categoryToDomain(category)
		if err != nil {
			return nil, fmt.Errorf("failed to create domain category: %w", err)
		}

		domainCategories = append(domainCategories, domainCategory)
	}

	return domainCategories, nil
}
