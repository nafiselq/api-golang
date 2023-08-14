package appcontext

import (
	"errors"

	"github.com/e-ziswaf/eziswaf-api/config"
	"github.com/e-ziswaf/eziswaf-api/internal/app/driver"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

const (
	// DBDialectMysql rdbms dialect name for MySQL
	DBDialectMysql = "mysql"

	// DBDialectPostgres rdbms dialect name for PostgreSQL
	DBDialectPostgres = "postgres"
)

// AppContext the app context struct
type AppContext struct {
	config config.Provider
}

// NewAppContext initiate appcontext object
func NewAppContext(config config.Provider) *AppContext {
	return &AppContext{
		config: config,
	}
}

// GetDBInstance getting gorp instance, param: dbType can be "mysql" or "postgre"
func (a *AppContext) GetDBInstance(dbType string) (*gorm.DB, error) {
	var gormDbMap *gorm.DB
	var err error
	switch dbType {
	case DBDialectMysql:
		dbOption := a.getMysqlOption()
		gormDbMap, err = driver.NewMysqlDatabase(dbOption)
	case DBDialectPostgres:
		dbOption := a.getPostgreOption()
		gormDbMap, err = driver.NewPostgreDatabase(dbOption)
	default:
		err = errors.New("Error get db instance, unknown db type")
	}

	return gormDbMap, err
}

func (a *AppContext) getMysqlOption() driver.DBMysqlOption {
	return driver.DBMysqlOption{
		Host:                 a.config.GetString("mysql.host"),
		Port:                 a.config.GetInt("mysql.port"),
		Username:             a.config.GetString("mysql.username"),
		Password:             a.config.GetString("mysql.password"),
		DBName:               a.config.GetString("mysql.name"),
		AdditionalParameters: a.config.GetString("mysql.additional_parameters"),
		MaxOpenConns:         a.config.GetInt("mysql.conn_open_max"),
		MaxIdleConns:         a.config.GetInt("mysql.conn_idle_max"),
		ConnMaxLifetime:      a.config.GetDuration("mysql.conn_lifetime_max"),
	}
}

func (a *AppContext) getPostgreOption() driver.DBPostgreOption {
	return driver.DBPostgreOption{
		Host:        a.config.GetString("postgre.host"),
		Port:        a.config.GetInt("postgre.port"),
		Username:    a.config.GetString("postgre.username"),
		Password:    a.config.GetString("postgre.password"),
		DBName:      a.config.GetString("postgre.name"),
		MaxPoolSize: a.config.GetInt("postgre.pool_size"),
	}
}

// GetCachePool get cache pool connection
func (a *AppContext) GetCachePool() *redis.Pool {
	return driver.NewCache(a.getCacheOption())
}

func (a *AppContext) getCacheOption() driver.CacheOption {
	return driver.CacheOption{
		Host:               a.config.GetString("cache.host"),
		Port:               a.config.GetInt("cache.port"),
		Namespace:          a.config.GetString("cache.namespace"),
		Password:           a.config.GetString("cache.password"),
		DialConnectTimeout: a.config.GetDuration("cache.dial_connect_timeout"),
		ReadTimeout:        a.config.GetDuration("cache.read_timeout"),
		WriteTimeout:       a.config.GetDuration("cache.write_timeout"),
		IdleTimeout:        a.config.GetDuration("cache.idle_timeout"),
		MaxConnLifetime:    a.config.GetDuration("cache.conn_lifetime_max"),
		MaxIdle:            a.config.GetInt("cache.conn_idle_max"),
		MaxActive:          a.config.GetInt("cache.conn_active_max"),
		Wait:               a.config.GetBool("cache.is_wait"),
	}
}
