package services

import (
	"context"

	"github.com/northwindman/book-shop/internal/app/domain"
)

type UserRepository interface {
	GetUser(ctx context.Context, username string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByID(ctx context.Context, id int) (domain.User, error)
}

type BookRepository interface {
	GetBook(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context, categoryIDs []int, limit, offset int) ([]domain.Book, error)
	CreateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	DeleteBook(ctx context.Context, id int) error
}

type CategoryRepository interface {
	GetCategory(ctx context.Context, id int) (domain.Category, error)
	GetCategories(ctx context.Context) ([]domain.Category, error)
	CreateCategory(ctx context.Context, category domain.Category) (domain.Category, error)
	UpdateCategory(ctx context.Context, category domain.Category) (domain.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

type CartRepository interface {
	GetCart(ctx context.Context, userID int) (domain.Cart, error)
	DeleteCart(ctx context.Context, userID int) error
	UpdateCartAndStocks(ctx context.Context, cart domain.Cart) error
	CheckStocks(ctx context.Context, cart domain.Cart) (bool, error)
}
