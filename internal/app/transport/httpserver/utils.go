package httpserver

import (
	"context"

	"github.com/northwindman/book-shop/internal/app/domain"
)

func toResponseBook(book domain.Book) BookResponse {
	return BookResponse{
		ID:         book.ID(),
		Title:      book.Title(),
		Year:       book.Year(),
		Author:     book.Author(),
		Price:      book.Price(),
		Stock:      book.Stock(),
		CategoryID: book.CategoryID(),
	}
}

func toResponseCategory(category domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID(),
		Name: category.Name(),
	}
}

func toDomainBook(bookRequest BookRequest) (domain.Book, error) {
	return domain.NewBook(domain.NewBookData{
		Title:      bookRequest.Title,
		Year:       bookRequest.Year,
		Author:     bookRequest.Author,
		Price:      bookRequest.Price,
		Stock:      bookRequest.Stock,
		CategoryID: bookRequest.CategoryID,
	})
}

func toDomainUser(username, password string) (domain.User, error) {
	return domain.NewUser(domain.NewUserData{
		Username: username,
		Password: password,
	})
}

func toDomainCart(userID int, cartRequest CartRequest) (domain.Cart, error) {
	return domain.NewCart(domain.NewCartData{
		UserID:  userID,
		BookIDs: cartRequest.BookIDs,
	})
}

func toResponseCart(cart domain.Cart) CartResponse {
	return CartResponse{
		BookIDs: cart.BookIDs(),
	}
}

func getUserFromContext(ctx context.Context) (domain.User, error) {
	contextUser := ctx.Value(ContextUserKey)
	if contextUser == nil {
		return domain.User{}, domain.ErrNoUserInContext
	}
	user, ok := contextUser.(domain.User)
	if !ok {
		return domain.User{}, domain.ErrNoUserInContext
	}
	return user, nil
}
