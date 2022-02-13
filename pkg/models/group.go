package models

type Group struct {
	Model
	Name          string `gorm:"column:name" json:"name,omitempty"`
	IsSystemGroup bool   `gorm:"column:is_system_group" json:"-"`
	Users         []User `gorm:"many2many:user_groups;"`
}
