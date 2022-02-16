package models

import (
	"github.com/jackc/pgtype"
)

const (
	SeeUsers          string = "see:users"
	CreateUser        string = "create:user"
	UpdateUser        string = "update:user"
	UpdateAccount     string = "update:account"
	UpdatePassword    string = "update:password"
	UpdateOwnPassword string = "update:own:password"
	DeleteUser        string = "delete:user"

	SeeServers         string = "see:servers"
	SeeServersByGroups string = "see:server:by:groups"

	SeeGroups   string = "see:groups"
	CreateGroup string = "create:group"
	UpdateGroup string = "update:group"
	DeleteGroup string = "delete:group"

	SeeSystemGroups   string = "see:system:groups"
	CreateSystemGroup string = "create:system:group"
	UpdateSystemGroup string = "update:system:group"
	DeleteSystemGroup string = "delete:system:group"

	SeeRoles   string = "see:roles"
	CreateRole string = "create:role"
	UpdateRole string = "update:role"
	DeleteRole string = "delete:role"
)

type Role struct {
	Model

	Name        string           `gorm:"column:name" json:"name,omitempty"`
	Permissions pgtype.TextArray `gorm:"type:text[];column:permissions" json:"permissions,omitempty"`

	Users []User `json:"-"`
}

func NewRole(name string, perm []string) Role {
	mapped := make([]pgtype.Text, 0, len(perm))

	for _, p := range perm {
		mapped = append(mapped, pgtype.Text{
			String: p,
			Status: pgtype.Present,
		})
	}

	permissions := pgtype.TextArray{
		Elements:   mapped,
		Status:     pgtype.Present,
		Dimensions: []pgtype.ArrayDimension{{Length: int32(len(perm))}},
	}

	return Role{
		Name:        name,
		Permissions: permissions,
	}
}

func GetDefaultRoles() []Role {
	return []Role{
		NewRole("Administrator", []string{
			// Users
			SeeUsers,
			CreateUser,
			UpdateUser,
			UpdateOwnPassword,
			UpdatePassword,
			DeleteUser,
			UpdateAccount,

			// Servers
			SeeServers,

			// Groups
			SeeGroups,
			CreateGroup,
			UpdateGroup,
			DeleteGroup,

			// System Groups
			SeeSystemGroups,
			CreateSystemGroup,
			UpdateSystemGroup,
			DeleteSystemGroup,

			// Roles
			SeeRoles,
			CreateRole,
			UpdateRole,
			DeleteRole,
		}),
		NewRole("User", []string{
			UpdateOwnPassword,
			UpdateAccount,
			SeeServersByGroups,
		}),
	}
}
