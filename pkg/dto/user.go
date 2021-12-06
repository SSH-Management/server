package dto

type User struct {
	Name         string   `json:"name" conform:"trim" validate:"required,max=50"`
	Surname      string   `json:"surname" conform:"trim" validate:"required,max=50"`
	Username     string   `json:"username" conform:"trim" validate:"required,max=50"`
	Email        string   `json:"email" conform:"trim,email" validate:"required,email,max=150"`
	Password     string   `json:"password" conform:"trim" validate:"required,min=6,max=50"`
	Shell        string   `json:"shell" conform:"trim" validate:"required,max=50"`
	Role         string   `json:"role" conform:"trim" validate:"required,max=50"`
	Groups       []string `json:"groups"`
	SystemGroups []string `json:"system_groups"  validate:"required"`
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetSurname() string {
	return u.Surname
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetShell() string {
	return u.Shell
}

func (u User) GetRole() string {
	return u.Role
}

func (u User) GetGroups() []string {
	return u.Groups
}

func (u User) GetSystemGroups() []string {
	return u.SystemGroups
}

func (u User) GetPlainTextPassword() string {
	return u.Password
}

func (u User) GetDefaultShell() string {
	return u.Shell
}

func (u User) GetUsername() string {
	return u.Username
}

func (u User) WithPassword(password string) User {
	return User{
		Name:         u.Name,
		Surname:      u.Surname,
		Username:     u.Username,
		Email:        u.Email,
		Password:     password,
		Shell:        u.Shell,
		Role:         u.Role,
		Groups:       u.Groups,
		SystemGroups: u.SystemGroups,
	}
}
