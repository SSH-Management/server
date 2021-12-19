package dto

type CreateUser struct {
	User         User   `json:"user,omitempty" validate:"required"`
	PublicSSHKey string `json:"public_key,omitempty" conform:"trim" validate:"required"`
}

// func (c CreateUser) GetUser() sdk.User {
// 	return c.User
// }

// func (c CreateUser) GetPublicKey() string {
// 	return c.PublicSSHKey
// }
