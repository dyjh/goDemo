package datasource

import (
	"demo/config"
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

var redisPool *redis.Pool

// 初始化redis
func InitRedis() (conn redis.Conn,err error) {
	redisConfig := config.InitConfig().Redis
	redisPool = &redis.Pool{
		MaxIdle:     0,
		MaxActive:   10,
		IdleTimeout: time.Duration(30) * time.Minute,
		Dial: func() (redis.Conn, error) {
			if redisConfig.DB == 0 {
				return redis.Dial("tcp", redisConfig.Addr + ":" + redisConfig.Port, redis.DialDatabase(redisConfig.DB))
			} else {
				return redis.Dial("tcp", redisConfig.Addr + ":" + redisConfig.Port)
			}
		},
	}

	conn = GetRedis()
	defer conn.Close()

	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		err = errors.New("redis connect failed.")
	}

	return conn, err
}

// 获取redis连接
func GetRedis() redis.Conn {
	return redisPool.Get()
}

// 关闭redis
func CloseRedis() {
	if redisPool != nil {
		_ = redisPool.Close()
	}
}

func Set(key, val string, ttl time.Duration) error {
	conn := GetRedis()
	defer conn.Close()

	r, err := redis.String(conn.Do("SET", key, val, "EX", ttl.Seconds()))

	if err != nil {
		return err
	}

	if r != "OK" {
		return errors.New("NOT OK")
	}

	return nil
}