// author by @xiaoyusilen

package route

import (
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis"
	"github.com/xiaoyusilen/geo-go/common/service"
	"github.com/xiaoyusilen/geo-go/config"
	mgo "gopkg.in/mgo.v2"
)

type RestApi struct {
	Router *gin.Engine

	// Configurations
	Config *config.Config

	// tile38
	Tile38 *redis.Client

	// mongodb
	Mongo *mgo.Session
}

func HandleRest(cfg *config.Config) *RestApi {

	r := &RestApi{

		Router: gin.Default(),

		Config: cfg,

		Tile38: service.NewTile38(cfg.Tile38Address),

		Mongo: service.NewMongo(cfg.RethinkAddress),
	}

	api := r.Router.Group("/api")
	{
		// todo: register route
		api.POST("/register", r.Register)
	}

	return r
}
