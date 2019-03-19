package api

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/jpz13/service/api/rest"
	"github.com/jpz13/service/core"
)

type Config struct {
	DummyService core.DummyService
	Logger       log.Logger
}

func New(config *Config) http.Handler {

	resources := rest.NewResources(&rest.Config{
		DummyService: config.DummyService,
		Logger:       config.Logger,
	})
	container := restful.NewContainer()

	// use the more performant CurlyRouter
	container.Router(restful.CurlyRouter{})

	// configure service analytics resources
	api := new(restful.WebService)
	api.Consumes(restful.MIME_JSON)
	api.Produces(restful.MIME_JSON)

	api.Route(api.POST("/dummy").To(resources.DummyFunction))
	container.Add(api)

	return container
}
