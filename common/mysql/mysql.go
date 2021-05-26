package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func Connect(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return gorm.Open(mysql.Open(defaultDSN), &gorm.Config{})
	}
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// ConnectWithConfig ConnectWithConfig
func ConnectWithConfig(dsn string, config *gorm.Config) (*gorm.DB, error) {
	if dsn == "" {
		return gorm.Open(mysql.Open(defaultDSN), config)
	}
	return gorm.Open(mysql.Open(dsn), config)
}
