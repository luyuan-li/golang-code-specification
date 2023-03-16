package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func NewUserRouter(g *gin.RouterGroup) {
	r := UserRouter{}

	g.GET("/list", r.UserList)
}

func (r *UserRouter) UserList(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}
