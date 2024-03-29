package service

// IHealthCheckService interface for health check service
type IHealthCheckService interface {
	HealthCheckDbMysql() (err error)
	HealthCheckDbPostgres() (err error)
	HealthCheckDbCache() (err error)
}

type healthCheck struct {
	opt Option
}

// NewHealthCheck create health check service instance with option as param
func NewHealthCheck(opt Option) IHealthCheckService {
	return &healthCheck{
		opt: opt,
	}
}

func (h *healthCheck) HealthCheckDbMysql() (err error) {
	dbObj, err := h.opt.DbMysql.DB()
	if err != nil {
		return
	}

	err = dbObj.Ping()
	if err != nil {
		// TODO: add logger (fatal err)
		// TODO: mapping err to db error
	}
	return
}

func (h *healthCheck) HealthCheckDbPostgres() (err error) {
	dbObj, err := h.opt.DbPostgre.DB()
	if err != nil {
		return
	}

	err = dbObj.Ping()
	if err != nil {
		// TODO: add logger (fatal err)
		// TODO: mapping err to db error
	}
	return
}

func (h *healthCheck) HealthCheckDbCache() (err error) {
	cacheConn := h.opt.CachePool.Get()
	_, err = cacheConn.Do("PING")
	if err != nil {
		// TODO: add logger (fatal err)
		// TODO: mapping err to cache err
		return
	}
	defer cacheConn.Close()

	return nil
}
