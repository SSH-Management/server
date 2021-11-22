package mysql

import (
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-sql-driver/mysql"

	"github.com/SSH-Management/server/pkg/db/config"
	"github.com/SSH-Management/server/pkg/db/connector"
)

type (
	mySql struct{}
)

func New() connector.Interface {
	return mySql{}
}

func (mySql) Connect(v config.Config) (gorm.Dialector, error) {
	config := mysql.NewConfig()

	config.Collation = v.Collation
	config.User = v.Username
	config.Passwd = v.Password
	config.DBName = v.Database
	config.ParseTime = true
	config.MultiStatements = true
	config.CheckConnLiveness = true
	config.Net = "tcp"
	config.Addr = v.Host

	connStr := config.FormatDSN()

	return gormmysql.Open(connStr), nil
}
