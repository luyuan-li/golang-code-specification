package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	GitTag  = ""
	GitHash = ""
)

type VersionRouter struct {
}

func NewVersionRouter(g *gin.RouterGroup) {
	r := VersionRouter{}

	g.GET("", r.GetVersion)
}

//
// GetVersion
//  @Description: 获取项目版本号
//  @receiver r
//  @param c
//
func (r *VersionRouter) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"git_tag":  GitTag,
		"git_hash": GitHash,
	})
}
