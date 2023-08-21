package driver

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // defines mysql driver used
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBMysqlOption options for mysql connection
type DBMysqlOption struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	DBName               string
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

// NewMysqlDatabase return gorm dbmap object with MySQL options param
func NewMysqlDatabase(option DBMysqlOption) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", option.Username, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	mysqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	mysqlDB.SetConnMaxLifetime(option.ConnMaxLifetime)
	mysqlDB.SetMaxIdleConns(option.MaxIdleConns)
	mysqlDB.SetMaxOpenConns(option.MaxOpenConns)

	err = mysqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
