package commons

import (
	"github.com/e-ziswaf/eziswaf-api/config"
	"github.com/e-ziswaf/eziswaf-api/internal/app/appcontext"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// Options common option for all object that needed
type Options struct {
	AppCtx         *appcontext.AppContext
	ProviderConfig config.Provider
	AppConfig      config.AppConfig
	DbMysql        *gorm.DB
	DbPostgre      *gorm.DB
	CachePool      *redis.Pool
}
