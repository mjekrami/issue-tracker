package dtos

type (
	User struct {
		ID        int
		FirstName string `json:"firstname" required:"true"`
		LastName  string `json:"lastname" required:"true"`
	}

	UserDTO struct {
		FirstName string
		LastName  string
	}
)
