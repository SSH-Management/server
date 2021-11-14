package models

type Role struct {
	Model
	Name string

	Users []User
}
