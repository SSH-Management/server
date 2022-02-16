package models

import (
	"database/sql"
	"database/sql/driver"
)

type (
	ServerStatus string
	Server       struct {
		Model
		Name            string         `gorm:"column:name" json:"name,omitempty"`
		IpAddress       string         `gorm:"column:ip" json:"ip,omitempty"`
		Status          ServerStatus   `gorm:"status" json:"status,omitempty"`
		PublicIpAddress sql.NullString `gorm:"column:public_ip" json:"public_ip,omitempty"`
		GroupID         uint64         `gorm:"column:group_id" json:"group_id,omitempty"`
		Group           Group          `json:"-"`
	}
)

const (
	ServerStatusUnknown    ServerStatus = "unknown"
	ServerStatusOk         ServerStatus = "ok"
	ServerStatusNotServing ServerStatus = "not_serving"
)

func (s *ServerStatus) Scan(data interface{}) error {
	*s = ServerStatus(data.([]byte))
	return nil
}

func (s ServerStatus) Value() (driver.Value, error) {
	return string(s), nil
}
