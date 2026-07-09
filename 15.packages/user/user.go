package user

// User is exported, so it can be used by other packages.
type User struct {
	Email string
	Name  string
	Phone string
	ID    int
}
