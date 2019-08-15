package api

import (
	"log"
	"net/http"

	"github.com/JPZ13/service/api/rest"
	"github.com/JPZ13/service/core"
	restful "github.com/emicklei/go-restful"
)

// Config holds api configuration
type Config struct {
	DummyService core.Service
	Logger       log.Logger
}

// New inits an api
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
	api.Path("/api/blog/v1")
	api.Consumes(restful.MIME_JSON)
	api.Produces(restful.MIME_JSON)

	api.Route(api.POST("/authors").To(resources.CreateAuthor))
	container.Add(api)

	return container
}
