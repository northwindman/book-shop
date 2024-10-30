//go:generate mockery

package httpserver

import (
	"context"

	"github.com/northwindman/book-shop/internal/app/domain"
)

// UserService is a user service
type UserService interface {
	GetUser(ctx context.Context, username string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByID(ctx context.Context, id int) (domain.User, error)
}

// TokenService is a token service
type TokenService interface {
	GenerateToken(user domain.User) (string, error)
	GetUser(token string) (domain.User, error)
}

// BookService is a book service
type BookService interface {
	GetBook(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context, categoryIDs []int, limit, offset int) ([]domain.Book, error)
	CreateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	DeleteBook(ctx context.Context, id int) error
}

// CategoryService is a category service
type CategoryService interface {
	GetCategory(ctx context.Context, id int) (domain.Category, error)
	GetCategories(ctx context.Context) ([]domain.Category, error)
	CreateCategory(ctx context.Context, category domain.Category) (domain.Category, error)
	UpdateCategory(ctx context.Context, category domain.Category) (domain.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

type CartService interface {
	UpdateCartAndStocks(ctx context.Context, cart domain.Cart) (domain.Cart, error)
	Checkout(ctx context.Context, userID int) error
}
