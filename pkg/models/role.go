package models

type Role struct {
	Model
	Name string `gorm:"column:name" json:"name,omitempty"`

	Users []User
}
