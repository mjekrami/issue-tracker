package models

// Privelege enum
type privilege int

const (
	User privilege = iota
	SuperUser
	Admin
)

type (
	UserCreate struct {
		Name      string    `json:"name" required:"true"`
		LastName  string    `json:"last_name" required:"true"`
		Privilege privilege `json:"privilege" required:"true"`
	}
	UserUpdate struct {
		ID string `json:"id" required:"true"`
	}
	UserResponse struct {
		*UserCreate
		JoinDate string `json:"join_date" required:"true"`
	}
)
