package router

import (
	"log"

	"github.com/golang-code-specification/internal/app/global"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	level := global.Config.App.LogLevel
	if level == "pro" || level == "prod" || level == "release" || level == "info" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	// log
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter)

	return r
}
