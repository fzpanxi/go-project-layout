package server

import (
	v1 "ddd_demo/api/ddd_demo/v1"
	"ddd_demo/internal/conf"
	application "ddd_demo/internal/usecase/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHttpServer)

func NewHttpServer(c *conf.Conf, uc *application.UserUseCase) *http.Server {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	server := &http.Server{
		Addr:           c.Server.Http.Addr,
		Handler:        v1.NewUserServiceHandler(uc, router),
		ReadTimeout:    c.Server.Http.Timeout,
		WriteTimeout:   c.Server.Http.Timeout,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}
