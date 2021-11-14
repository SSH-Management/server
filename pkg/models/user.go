package models

type User struct {
	Model
	Name         string `gorm:"column:name" json:"name,omitempty"`
	Surname      string `gorm:"column:surname" json:"surname,omitempty"`
	Username     string `gorm:"column:username" json:"username,omitempty"`
	Email        string `gorm:"column:email" json:"email,omitempty"`
	Password     string `gorm:"column:password" json:"password,omitempty"`
	Shell        string `gorm:"column:shell" json:"shell,omitempty"`
	PublicSSHKey string `gorm:"column:public_ssh_key" json:"public_ssh_key"`

	RoleID uint64 `gorm:"column:role_id" json:"-"`
	Role   Role   `json:"-"`

	Groups []Group `gorm:"many2many:user_groups;" json:"groups,omitempty"`
}
