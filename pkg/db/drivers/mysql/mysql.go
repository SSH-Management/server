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

func FormatDSN(v config.Config, withDb bool) string {
	mysqlConfig := mysql.NewConfig()

	mysqlConfig.Collation = v.Collation
	mysqlConfig.User = v.Username
	mysqlConfig.Passwd = v.Password

	if withDb {
		mysqlConfig.DBName = v.Database
	}

	mysqlConfig.ParseTime = true
	mysqlConfig.MultiStatements = true
	mysqlConfig.CheckConnLiveness = true
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = v.Host

	return mysqlConfig.FormatDSN()
}

func (mySql) Connect(v config.Config) (gorm.Dialector, error) {
	return gormmysql.Open(FormatDSN(v, true)), nil
}
