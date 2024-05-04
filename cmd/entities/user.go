package entities

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Metadata  *UserMetadata
}

// User metadata stores some data the is beyond use of application scope
type UserMetadata struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
