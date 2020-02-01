package datasource

import (
	"demo/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

/**
 * 返回Redis实例
 */
func NewRedis() *redis.Database {

	var database *redis.Database

	//项目配置
	appConfig := config.InitConfig()
	if appConfig != nil {
		iris.New().Logger().Info(" hello ")
		rd := appConfig.Redis
		iris.New().Logger().Info(rd)
		database = redis.New(redis.Config{
			Network:     rd.NetWork,
			Addr:        rd.Addr + ":" + rd.Port,
			Password:    rd.Password,
			Database:    "",
			MaxActive:   10,
			Timeout: redis.DefaultRedisTimeout,
			Prefix:      rd.Prefix,
		})
	} else {
		iris.New().Logger().Info(" hello  error ")
	}
	return database
}
