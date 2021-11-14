package models

type (
	Group struct {
		Model
		Name  string
		Users []User `gorm:"many2many:user_groups;"`
	}
)
