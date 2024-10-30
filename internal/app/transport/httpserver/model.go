package httpserver

import (
	"fmt"

	"github.com/northwindman/book-shop/internal/app/domain"
)

type BookRequest struct {
	Title      string `json:"title"`
	Year       int    `json:"year"`
	Author     string `json:"author"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

func (r *BookRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("%w: title", domain.ErrRequired)
	}
	if r.Year <= 0 {
		return fmt.Errorf("%w: year", domain.ErrNegative)
	}
	if r.Author == "" {
		return fmt.Errorf("%w: author", domain.ErrRequired)
	}
	if r.Price <= 0 {
		return fmt.Errorf("%w: price", domain.ErrNegative)
	}
	if r.CategoryID == 0 {
		return fmt.Errorf("%w: category_id", domain.ErrRequired)
	}
	return nil
}

type BookResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Year       int    `json:"year"`
	Author     string `json:"author"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

type CategoryRequest struct {
	Name string `json:"name"`
}

func (r *CategoryRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("%w: name", domain.ErrRequired)
	}
	return nil
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AuthRequest) Validate() error {
	if r.Username == "" {
		return fmt.Errorf("%w: username", domain.ErrRequired)
	}
	if r.Password == "" {
		return fmt.Errorf("%w: password", domain.ErrRequired)
	}
	return nil
}

type CartRequest struct {
	BookIDs []int `json:"book_ids"`
}

type CartResponse struct {
	BookIDs []int `json:"book_ids"`
}
