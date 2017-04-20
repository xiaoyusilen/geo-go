// author by @xiaoyusilen

package service

import (
	log "github.com/Sirupsen/logrus"
	redis "gopkg.in/redis.v5"
	"testing"
)

func TestNewTile38Service(t *testing.T) {

	// TODO: Add other test example
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:9851",
	})
	cmd := redis.NewStringCmd("SET", "fleet", "truck", "POINT", 33.32, 115.423)
	client.Process(cmd)
	v, _ := cmd.Result()
	log.Println(v)
	cmd1 := redis.NewStringCmd("GET", "fleet", "truck")
	client.Process(cmd1)
	v1, _ := cmd1.Result()
	log.Println(v1)
}
