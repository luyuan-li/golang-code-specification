package rest

import (
	"net/http"

	"github.com/golang-code-specification/internal/app/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func NewUserRouter(g *gin.RouterGroup) {
	r := UserRouter{}

	g.GET("/list", r.UserList)
}

//
// UserList TODO 此处的返回值，根据项目返回响应的结构
//  @Description: 查询用户列表
//  @receiver r
//  @param c
//
func (r *UserRouter) UserList(c *gin.Context) {
	resp, err := service.UserService.UserList()
	if err != nil {
		c.JSON(http.StatusOK, "")
		return
	}

	c.JSON(http.StatusOK, resp)
}
