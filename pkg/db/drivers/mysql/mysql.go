package mysql

import (
	"gorm.io/gorm"

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

	// config := mysql.NewConfig()

	// config.Collation = v.Collation
	// config.User = v.Username
	// config.Passwd = v.Password
	// config.DBName = v.Database
	// config.ParseTime = true
	// config.MultiStatements = true
	// config.CheckConnLiveness = true
	// config.Addr = fmt.Sprintf("tcp(%s)", v.Host)

	// connStr := config.FormatDSN()

	// db, err := sql.Open("mysql", connStr)

	// if err != nil {
	// 	return nil, err
	// }

	// db.SetConnMaxIdleTime(v.MaxIdleTime)
	// db.SetConnMaxLifetime(v.ConnMaxLifetime)
	// db.SetMaxIdleConns(v.ConnMaxIdle)
	// db.SetMaxOpenConns(v.ConnMaxOpen)

	// return db, nil

	return nil, nil
}
