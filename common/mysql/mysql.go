package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DSN is data source name
	DSN string
)

func Init(dsn string) {
	DSN = dsn
}

func conn() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(DSN), &gorm.Config{})
}
