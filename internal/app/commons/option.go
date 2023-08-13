package commons

import (
	"github.com/gomodule/redigo/redis"
	"github.com/e-ziswaf/eziswaf-api/config"
	"github.com/e-ziswaf/eziswaf-api/internal/app/appcontext"
	"gopkg.in/gorp.v2"
)

// Options common option for all object that needed
type Options struct {
	AppCtx         *appcontext.AppContext
	ProviderConfig config.Provider
	AppConfig      config.AppConfig
	DbMysql        *gorp.DbMap
	DbPostgre      *gorp.DbMap
	CachePool      *redis.Pool
}
