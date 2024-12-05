package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/landrade/gqlgen-cache-control-plugin/cache"
)

// NewPlaygroundController
func NewPlaygroundController(config *Config, e *echo.Echo) {
	if p := config.Graphql.Playground; p.Enabled {
		e.GET(p.Path, echo.WrapHandler(playground.Handler("Playground", config.Graphql.Playground.QueryPath)))
	}
}

// NewQueryController
func NewQueryController(config *Config, server *handler.Server, e *echo.Echo) {
	e.Any(config.Graphql.Query.Path, echo.WrapHandler(cache.Middleware(server)))
}
