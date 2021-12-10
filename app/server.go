package app

import (
	"ecommerce/pkg/app"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server struct {
	config *config
	mode *app.Mode

	//handler http.Handler
	handler *gin.Engine
	core    *http.Server
}

func newServer(config *config, mode *app.Mode) *server {
	server := new(server)
	server.config = config
	server.mode = mode

	gin.SetMode(mode.String())
	server.handler = gin.New()

	server.initRouter()
	return server
}

func (s *server) run() {
	s.core = &http.Server{
		Addr:    s.config.Server.Port,
		Handler: s.handler,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	go func() {
		//logger.WithContext(ctx).Printf("HTTP server is running at %s.", s.config.Port)
		err := s.core.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			//logger.WithContext(ctx).Errorf("server run error: %s.", err.Error())
			//log.Fatal("server run error:", err)
		}
	}()
}

