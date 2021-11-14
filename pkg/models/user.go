package models

type User struct {
	Model        `json:"model,omitempty"`
	Name         string `json:"name,omitempty"`
	Surname      string `json:"surname,omitempty"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	Shell        string `json:"shell,omitempty"`
	PublicSSHKey string `json:"public_ssh_key"`

	RoleID uint64 `json:"-"`
	Role   Role   `json:"-"`

	Groups []Group `gorm:"many2many:user_groups;" json:"groups,omitempty"`
}
