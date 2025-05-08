package validators

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=50"`
	LastName  string `json:"last_name" validate:"required,min=3,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=40"`
	Age       uint8  `json:"age" validate:"required,min=18,max=130"`
}
