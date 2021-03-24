package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// defaultDSN is the default data source name
	defaultDSN string
)

func Init(dsn string) {
	defaultDSN = dsn
}

func Conn(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return gorm.Open(mysql.Open(defaultDSN), &gorm.Config{})
	}
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
