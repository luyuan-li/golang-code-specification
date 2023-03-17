package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-code-specification/internal/app/global"
)

func NewRouter() *gin.Engine {
	level := global.Config.App.LogLevel
	if level == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	// log
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter)

	return r
}
