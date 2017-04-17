// author by @xiaoyusilen

package service

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

var Cache CacheService

type CacheService struct {
	pool redis.Pool
}

func NewTile38Service(url string, replica, maxIdleConnection int, idleTimeout, connectTimeout, readTimeout, writeTimeout time.Duration) *CacheService {
	cache := &CacheService{}
	if err := cache.Init(url, replica, maxIdleConnection, idleTimeout, connectTimeout, readTimeout, writeTimeout); err != nil {
		log.Panic("Init tile38 failed.")
	}

	log.Info("Init tile38 success.")

	return cache
}

// 初始化连接
func (cache *CacheService) Init(ip string, dbNumber, maxIdleConnection int, idleTimeout, connectTimeout, readTimeout, writeTimeout time.Duration) error {

	cache.pool = redis.Pool{
		MaxIdle:     maxIdleConnection,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := InitRedis("tcp", ip, connectTimeout, readTimeout, writeTimeout, dbNumber)
			if err != nil {
				log.Error("tile38 pool init connection:", err, conn)
				return nil, err
			}
			log.Debug("tile38 pool init connection sucess!")
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	log.Error("connect tile38 pool success!", dbNumber)

	return nil
}

func InitRedis(network, address string, connectTimeout, readTimeout, writeTimeout time.Duration, db int) (redis.Conn, error) {
	return redis.Dial(network, address,
		redis.DialConnectTimeout(connectTimeout),
		redis.DialReadTimeout(readTimeout),
		redis.DialWriteTimeout(writeTimeout),
		redis.DialDatabase(db))
}
