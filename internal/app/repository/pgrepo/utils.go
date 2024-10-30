package pgrepo

import (
	"github.com/northwindman/book-shop/internal/app/domain"
	"github.com/northwindman/book-shop/internal/app/repository/models"
)

func domainToBook(book domain.Book) models.Book {
	return models.Book{
		ID:         book.ID(),
		Title:      book.Title(),
		Year:       book.Year(),
		Author:     book.Author(),
		Price:      book.Price(),
		Stock:      book.Stock(),
		CategoryID: book.CategoryID(),
	}
}

func bookToDomain(book models.Book) (domain.Book, error) {
	return domain.NewBook(domain.NewBookData{
		ID:         book.ID,
		Title:      book.Title,
		Year:       book.Year,
		Author:     book.Author,
		Price:      book.Price,
		Stock:      book.Stock,
		CategoryID: book.CategoryID,
	})
}

func domainToCategory(category domain.Category) models.Category {
	return models.Category{
		ID:   category.ID(),
		Name: category.Name(),
	}
}

func categoryToDomain(category models.Category) (domain.Category, error) {
	return domain.NewCategory(domain.NewCategoryData{
		ID:   category.ID,
		Name: category.Name,
	})
}

func domainToUser(user domain.User) models.User {
	return models.User{
		ID:       user.ID(),
		Username: user.Username(),
		Password: user.Password(),
		Admin:    user.Admin(),
	}
}

func userToDomain(user models.User) (domain.User, error) {
	return domain.NewUser(domain.NewUserData{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Admin:    user.Admin,
	})
}

func domainToCart(cart domain.Cart) models.Cart {
	return models.Cart{
		UserID:  cart.UserID(),
		BookIDs: cart.BookIDs(),
	}
}

func cartToDomain(cart models.Cart) (domain.Cart, error) {
	return domain.NewCart(domain.NewCartData{
		UserID:  cart.UserID,
		BookIDs: cart.BookIDs,
	})
}
