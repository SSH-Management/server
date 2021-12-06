package dto

type (
	Login struct {
		Email    string `json:"email,omitempty" conform:"trim,email" validate:"required,email,max=150"`
		Password string `json:"password,omitempty" conform:"trim" validate:"required,min=6,max=50"`
	}
)
