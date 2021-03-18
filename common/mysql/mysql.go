package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func conn(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
