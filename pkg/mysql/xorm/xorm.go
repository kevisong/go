package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	x "xorm.io/xorm"
)

var (
	// defaultDSN is the default data source name
	defaultDSN string
)

// Init Init
func Init(dsn string) {
	defaultDSN = dsn
}

// Connect Connect
func Connect(dsn string) (*x.Engine, error) {
	if dsn == "" {
		return x.NewEngine("mysql", defaultDSN)
	}
	return x.NewEngine("mysql", dsn)
}
