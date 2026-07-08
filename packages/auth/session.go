package auth

import "time"

// Session stores basic login session details.
type Session struct {
	Username  string
	CreatedAt time.Time
}

// NewSession creates a session value for the provided username.
func NewSession(username string) Session {
	return Session{
		Username:  username,
		CreatedAt: time.Now(),
	}
}
