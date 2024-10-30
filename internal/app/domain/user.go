package domain

// User is a domain User.
type User struct {
	id       int
	username string
	password string
	admin    bool
}

type NewUserData struct {
	ID       int
	Username string
	Password string
	Admin    bool
}

// NewUser creates a new user.
func NewUser(data NewUserData) (User, error) {
	return User{
		id:       data.ID,
		username: data.Username,
		password: data.Password,
		admin:    data.Admin,
	}, nil
}

// ID returns the user ID.
func (u User) ID() int {
	return u.id
}

// Username returns the user username.
func (u User) Username() string {
	return u.username
}

// Password returns the user password.
func (u User) Password() string {
	return u.password
}

// Admin returns the user admin flag.
func (u User) Admin() bool {
	return u.admin
}
