package models

type Role struct {
	Model

	Name string `gorm:"column:name" json:"name,omitempty"`
	// Permissions types.StringArray `gorm:"column:permissions" json:"permissions,omitempty"`

	Users []User `json:"-"`
}
