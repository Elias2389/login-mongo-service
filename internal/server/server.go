package server

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"login-mongo-service/config"
	"login-mongo-service/pkg/logger"
	"net/http"
	"time"
)

const (
	certFile       = "ssl/Server.crt"
	keyFile        = "ssl/Server.pem"
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server Struct
type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	mongo  *mongo.Client
	logger logger.Logger
}

// New server
func NewServer(cfg *config.Config, mongo *mongo.Client, logger logger.Logger) *Server {
	return &Server{echo: echo.New(), cfg: cfg, mongo: mongo, logger: logger}
}

// IniT server
func (s Server) RunServer() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
	if err := s.echo.StartServer(server); err != nil {
		s.logger.Fatal("Error starting server", err)
	}

	return nil
}
