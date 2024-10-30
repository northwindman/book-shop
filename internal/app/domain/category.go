package domain

// Category is a domain category.
type Category struct {
	id   int
	name string
}

type NewCategoryData struct {
	ID   int
	Name string
}

// NewCategory creates a new book.
func NewCategory(data NewCategoryData) (Category, error) {
	return Category{
		id:   data.ID,
		name: data.Name,
	}, nil
}

// ID returns the book ID.
func (b Category) ID() int {
	return b.id
}

// Name returns the book name.
func (b Category) Name() string {
	return b.name
}
