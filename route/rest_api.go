// author by @xiaoyusilen

package route

import (
	r "github.com/GoRethink/gorethink.git"
	log "github.com/Sirupsen/logrus"
	"github.com/ant0ine/go-json-rest.git/rest"
	"github.com/geo-go/common/service"
	"github.com/geo-go/config"
)

type RestApi struct {
	*rest.Api

	// Configurations
	Config *config.Config

	// tile38
	Cache *service.CacheService

	// rethinkdb
	Rethink *r.Session
}

func NewRestApi(cfg *config.Config) *RestApi {

	r := &RestApi{

		Api: rest.NewApi(),

		Config: cfg,

		Cache: service.NewTile38Service(cfg.Tile38Address,
			cfg.Tile38DBNumber,
			cfg.Tile38MaxIdleConnection,
			cfg.Tile38IdleTimeout,
			cfg.Tile38ConnectTimeout,
			cfg.Tile38ReadTimeout,
			cfg.Tile38WriteTimeout,
		),

		Rethink: service.NewRethinkdb(cfg.RethinkAddress),
	}

	r.SetApp(r.makeRouter())

	return r
}

func (api *RestApi) makeRouter() rest.App {

	routers := []*rest.Route{}

	routers = append(routers, []*rest.Route{
	// add rest api func
	}...)

	app, err := rest.MakeRouter(routers...)

	if err != nil {
		log.Panic(err)
	}

	return app
}
