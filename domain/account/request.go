package account

// Request body for account registration
type AccountRegisterRequest struct {
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

// Request body for account login
type AccountLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
