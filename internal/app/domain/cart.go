package domain

type Cart struct {
	userID  int
	bookIDs []int
}

type NewCartData struct {
	UserID  int
	BookIDs []int
}

func NewCart(data NewCartData) (Cart, error) {
	return Cart{
		userID:  data.UserID,
		bookIDs: data.BookIDs,
	}, nil
}

func (c Cart) UserID() int {
	return c.userID
}

func (c Cart) BookIDs() []int {
	return c.bookIDs
}

func (c Cart) Diff(old Cart) Cart {
	diff := Cart{
		userID:  c.userID,
		bookIDs: []int{},
	}

	oldCartBookIDs := map[int]struct{}{}
	for _, bookID := range old.bookIDs {
		oldCartBookIDs[bookID] = struct{}{}
	}

	for _, bookID := range c.bookIDs {
		if _, ok := oldCartBookIDs[bookID]; !ok {
			diff.bookIDs = append(diff.bookIDs, bookID)
		}
	}

	return diff
}

func (c Cart) HasBooks() bool {
	return len(c.bookIDs) > 0
}

func (c Cart) Equal(other Cart) bool {
	if c.userID != other.userID {
		return false
	}

	if len(c.bookIDs) != len(other.bookIDs) {
		return false
	}

	for i, bookID := range c.bookIDs {
		if bookID != other.bookIDs[i] {
			return false
		}
	}

	return true
}

func (c Cart) Join(other Cart) Cart {
	bookIDs := map[int]struct{}{}
	for _, bookID := range c.bookIDs {
		bookIDs[bookID] = struct{}{}
	}

	for _, bookID := range other.bookIDs {
		bookIDs[bookID] = struct{}{}
	}

	joined := Cart{
		userID:  c.userID,
		bookIDs: []int{},
	}

	for bookID := range bookIDs {
		joined.bookIDs = append(joined.bookIDs, bookID)
	}

	return joined
}
