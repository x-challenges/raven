package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/worker"
)

// ServerWorker interface.
type ServerWorker interface {
	worker.Worker
}

// ServerWorker interface implementation.
type serverWorker struct {
	logger *zap.Logger
	echo   *echo.Echo
	server *http.Server
}

// NewServerWorker construct.
func NewServerWorker(logger *zap.Logger, echo *echo.Echo, server *http.Server) ServerWorker {
	return &serverWorker{
		logger: logger,
		echo:   echo,
		server: server,
	}
}

// Run implements worker.Worker interface.
func (s *serverWorker) Run(_ context.Context) error {
	defer s.logger.Info("http server started")

	go func() {
		if err := s.echo.StartServer(s.server); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				s.logger.Fatal("http server start faield", zap.Error(err))
			}
		}
	}()
	return nil
}

// Shutdown implements worker.Worker interface.
func (s *serverWorker) Shutdown(ctx context.Context) error {
	defer s.logger.Info("http server stopped")
	return s.echo.Shutdown(ctx)
}
