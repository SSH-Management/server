package dto

type DeleteUser struct {
	Username string `json:"username" conform:"trim" validate:"required,max=50"`
}
