package models

import "database/sql"

type Server struct {
	Model

	Name            string
	IpAddress       string
	PublicIpAddress sql.NullString
	GroupID         uint64
	Group           Group
}
