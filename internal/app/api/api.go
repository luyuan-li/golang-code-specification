package api

import (
	"net/http"

	"github.com/golang-code-specification/internal/app/api/middleware"
	"github.com/golang-code-specification/internal/app/api/rest"
	"github.com/golang-code-specification/internal/app/api/router"
	"github.com/golang-code-specification/internal/app/api/server"
	"github.com/golang-code-specification/internal/app/config"
)

func RegisterApi(conf *config.App) server.Server {
	r := router.NewRouter()
	//TODO 此处仅距离Cors中间件，如果需要使用其他中间件，在下面加上即可
	r.Use(middleware.Cors())

	commonGroup := r.Group("")

	rest.NewUserRouter(commonGroup.Group("/user"))
	rest.NewVersionRouter(commonGroup.Group("/version"))

	srv := &http.Server{
		Addr:    conf.Address,
		Handler: r,
	}
	return server.NewServer(srv)

}
