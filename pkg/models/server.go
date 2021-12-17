package models

import "database/sql"

type Server struct {
	Model
	Name            string         `gorm:"column:name" json:"name,omitempty"`
	IpAddress       string         `gorm:"column:ip" json:"ip,omitempty"`
	Status          string         `gorm:"status" json:"status,omitempty"`
	PublicIpAddress sql.NullString `gorm:"column:public_ip" json:"public_ip,omitempty"`
	GroupID         uint64         `gorm:"column:group_id" json:"group_id,omitempty"`
	Group           Group
}
