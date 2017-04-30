// author by @xiaoyusilen

package service

import (
	log "github.com/Sirupsen/logrus"
	redis "github.com/go-redis/redis"
)

func NewTile38(url string) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr: url,
	})

	log.Debugf("Tile38 connect success.")

	return client
}
