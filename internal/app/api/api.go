package api

import (
	"net/http"

	"github.com/golang-code-specification/internal/app/api/router"

	"github.com/golang-code-specification/internal/app/config"

	"github.com/golang-code-specification/internal/app/api/middleware"
	"github.com/golang-code-specification/internal/app/api/rest"
	"github.com/golang-code-specification/internal/app/api/server"
)

func RegisterApi(conf *config.App) server.Server {
	r := router.NewRouter()
	r.Use(middleware.Cors())

	commonGroup := r.Group("")

	rest.NewUserRouter(commonGroup.Group("/user"))

	srv := &http.Server{
		Addr:    conf.Addr,
		Handler: r,
	}
	return server.NewServer(srv)

}
