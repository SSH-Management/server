package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

type (
	ServerStatus string
	Server       struct {
		Model
		Name            string         `gorm:"column:name" json:"name,omitempty"`
		IpAddress       string         `gorm:"column:ip" json:"ip,omitempty"`
		Status          ServerStatus   `gorm:"column:status" json:"status,omitempty"`
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
	status, ok := data.(string)

	if !ok {
		return errors.New("invalid data type in database")
	}

	*s = ServerStatus(status)
	return nil
}

func (s ServerStatus) Value() (driver.Value, error) {
	return string(s), nil
}
