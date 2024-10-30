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
	"github.com/uptrace/bun"
)

type BookRepo struct {
	db *pg.DB
}

func NewBookRepo(db *pg.DB) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (r BookRepo) GetBook(ctx context.Context, id int) (domain.Book, error) {
	if id == 0 {
		return domain.Book{}, fmt.Errorf("%w: id", domain.ErrRequired)
	}

	var book models.Book
	err := r.db.NewSelect().Model(&book).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Book{}, domain.ErrNotFound
		}
		return domain.Book{}, fmt.Errorf("failed to get a book: %w", err)
	}

	domainBook, err := bookToDomain(book)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to create domain book: %w", err)
	}

	return domainBook, nil
}

func (r BookRepo) CreateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	dbBook := domainToBook(book)

	var insertedBook models.Book
	err := r.db.NewInsert().Model(&dbBook).Returning("*").Scan(ctx, &insertedBook)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to insert a book: %w", err)
	}

	domainBook, err := bookToDomain(insertedBook)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to create domain book: %w", err)
	}

	return domainBook, nil

}

func (r BookRepo) UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	dbBook := domainToBook(book)
	dbBook.UpdatedAt = time.Now()

	var updatedBook models.Book
	err := r.db.NewUpdate().
		Model(&dbBook).
		Where("id = ?", dbBook.ID).
		ExcludeColumn("created_at", "stock").
		Returning("*").
		Scan(ctx, &updatedBook)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to update a book: %w", err)
	}

	domainBook, err := bookToDomain(updatedBook)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to create domain book: %w", err)
	}

	return domainBook, nil
}

func (r BookRepo) DeleteBook(ctx context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf("%w: id", domain.ErrRequired)
	}

	_, err := r.db.NewDelete().Model((*models.Book)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete a book: %w", err)
	}

	return nil
}

func (r BookRepo) GetBooks(ctx context.Context, categoryIDs []int, limit, offset int) ([]domain.Book, error) {
	var books []models.Book
	query := r.db.NewSelect().Model(&books)
	query.Where("stock > 0")
	if len(categoryIDs) > 0 {
		query.Where("category_id IN (?)", bun.In(categoryIDs))
	}
	if limit > 0 {
		query.Limit(limit)
	}
	if offset > 0 {
		query.Offset(offset)
	}
	query.Order("id")
	err := query.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}

	domainBooks := make([]domain.Book, len(books))
	for i, book := range books {
		domainBook, err := bookToDomain(book)
		if err != nil {
			return nil, fmt.Errorf("failed to create domain book: %w", err)
		}

		domainBooks[i] = domainBook
	}

	return domainBooks, nil
}
