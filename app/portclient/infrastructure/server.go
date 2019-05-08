package infrastructure

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// ServerInstance represents app HTTP server with helpers functions like start and shutdown
type ServerInstance struct {
	addr     string
	handlers *APIHandler

	httpServer *http.Server
}

// NewServer is just constructor for ServerInstance
func NewServer(addr string, handlers *APIHandler) *ServerInstance {
	s := &ServerInstance{
		addr:     addr,
		handlers: handlers,
	}

	return s
}

// Start will start http server and block execution on ListenAndServe call
func (s *ServerInstance) Start() {
	s.httpServer = &http.Server{Addr: s.addr, Handler: s.handlers.router}
	log.Infof("start to serve API on addr: %s", s.addr)

	err := s.httpServer.ListenAndServe()

	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("http Server stopped unexpected")
		s.Shutdown()
	} else {
		logrus.WithError(err).Info("http Server stopped")
	}
}

// Shutdown will gracefully stop server with 10 sec timeout
func (s *ServerInstance) Shutdown() {
	if s.httpServer != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logrus.WithError(err).Error("Failed to shutdown http server gracefully")
		} else {
			s.httpServer = nil
		}
	}
}
