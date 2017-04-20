// author by @xiaoyusilen

package route

import (
	r "github.com/GoRethink/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/xiaoyusilen/geo-go/common/service"
	"github.com/xiaoyusilen/geo-go/config"
	redis "gopkg.in/redis.v5"
)

type RestApi struct {
	Router *gin.Engine

	// Configurations
	Config *config.Config

	// tile38
	Tile38 *redis.Client

	// rethinkdb
	Rethink *r.Session
}

func HandleRest(cfg *config.Config) *RestApi {

	r := &RestApi{

		Router: gin.Default(),

		Config: cfg,

		Tile38: service.NewTile38(cfg.Tile38Address),

		Rethink: service.NewRethinkdb(cfg.RethinkAddress),
	}

	api := r.Router.Group("/api")
	{
		// todo: register route
		api.POST("/register", r.Register)
	}

	return r
}
