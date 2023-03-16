package app

import (
	"github.com/golang-code-specification/internal/app/api"
	"github.com/golang-code-specification/internal/app/config"
	"github.com/golang-code-specification/internal/app/global"
)

func Serve(conf *config.Config) {
	global.Config = conf

	server := api.RegisterApi(&conf.App)
	server.Start()
}
