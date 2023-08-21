package driver

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBPostgreOption options for postgre connection
type DBPostgreOption struct {
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	MaxPoolSize int
}

// NewPostgreDatabase return gorm dbmap object with postgre options param
func NewPostgreDatabase(option DBPostgreOption) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", option.Host, option.Username, option.Password, option.DBName, option.Port)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = pgsqlDB.Ping()
	if err != nil {
		return nil, err
	}

	pgsqlDB.SetMaxOpenConns(option.MaxPoolSize)

	return db, nil
}
