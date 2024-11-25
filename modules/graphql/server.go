package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/landrade/gqlgen-cache-control-plugin/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type NewServerParamsFx struct {
	fx.In
	Logger           *zap.Logger
	Config           *Config
	Schema           graphql.ExecutableSchema
	APQCache         *apqCacheAdapter
	ErrorPresenterFn errorPresenterFn
	RecoveryFn       recoveryFn
}

func NewServer(p NewServerParamsFx) *handler.Server {
	var (
		server = handler.New(p.Schema)
	)

	server.AddTransport(transport.Options{})
	server.AddTransport(transport.POST{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.MultipartForm{})
	server.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin:     func(_ *http.Request) bool { return true },
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		PingPongInterval:      p.Config.Graphql.Websocket.Ping,
		KeepAlivePingInterval: p.Config.Graphql.Websocket.KeepAlive,
	})

	server.Use(extension.FixedComplexityLimit(p.Config.Graphql.ComplexityLimit))

	if p.Config.Graphql.Tracing {
		server.Use(apollotracing.Tracer{})
	}

	if p.Config.Graphql.Intorspection {
		server.Use(extension.Introspection{})
	}

	if p.Config.Graphql.APQ.Enabled {
		server.Use(extension.AutomaticPersistedQuery{
			Cache: p.APQCache,
		})
	}

	// set custom error presenter
	server.SetErrorPresenter(p.ErrorPresenterFn)

	// recovery
	server.SetRecoverFunc(p.RecoveryFn)

	// cache control
	server.Use(cache.Extension{})

	return server
}
