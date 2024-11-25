package graphql

import (
	"path"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/landrade/gqlgen-cache-control-plugin/cache"
)

func NewPlaygroundController(config *Config, e *echo.Echo) {
	if p := config.Graphql.Playground; p.Enabled {

		var path = path.Join(
			"/1bb70d79-5756-4674-ae28-be0e258e664c", // TODO
			config.Graphql.Query.Path,
		)

		e.GET(p.Path, echo.WrapHandler(playground.Handler("Playground", path)))
	}
}

func NewQueryController(config *Config, server *handler.Server, e *echo.Echo) {
	var path = path.Join(
		"/:tenant",
		config.Graphql.Query.Path,
	)

	e.Any(path, echo.WrapHandler(cache.Middleware(server)))
}
