package models

type Group struct {
	Model
	Name  string `gorm:"column:name" json:"name,omitempty"`
	Users []User `gorm:"many2many:user_groups;"`
}
