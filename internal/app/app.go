package app

import (
	"github.com/golang-code-specification/internal/app/api"
	"github.com/golang-code-specification/internal/app/config"
	"github.com/golang-code-specification/internal/app/global"
	"github.com/golang-code-specification/internal/app/pkg"
	"github.com/golang-code-specification/internal/app/repository"
	"github.com/golang-code-specification/internal/app/service"
)

func Serve(conf *config.Config) {
	global.Config = conf

	mysqlDB := pkg.InitMysqlDB(conf.Mysql)

	repository.Init(mysqlDB)
	service.Init()

	server := api.RegisterApi(&conf.App)
	server.Start()
}
