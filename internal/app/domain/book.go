package domain

// Book is a domain book.
type Book struct {
	id         int
	title      string
	year       int
	author     string
	price      int
	stock      int
	categoryID int
}

type NewBookData struct {
	ID         int
	Title      string
	Year       int
	Author     string
	Price      int
	Stock      int
	CategoryID int
}

// NewBook creates a new book.
func NewBook(data NewBookData) (Book, error) {
	return Book{
		id:         data.ID,
		title:      data.Title,
		year:       data.Year,
		author:     data.Author,
		price:      data.Price,
		stock:      data.Stock,
		categoryID: data.CategoryID,
	}, nil
}

// ID returns the book ID.
func (b Book) ID() int {
	return b.id
}

// Title returns the book title.
func (b Book) Title() string {
	return b.title
}

// Year returns the book year.
func (b Book) Year() int {
	return b.year
}

// Author returns the book author.
func (b Book) Author() string {
	return b.author
}

// Price returns the book price.
func (b Book) Price() int {
	return b.price
}

// Stock returns the book stock.
func (b Book) Stock() int {
	return b.stock
}

// CategoryID returns the book category ID.
func (b Book) CategoryID() int {
	return b.categoryID
}
