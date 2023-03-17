package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-code-specification/internal/app/model/vo/req"
	"github.com/golang-code-specification/internal/app/service"
)

type UserRouter struct {
}

func NewUserRouter(g *gin.RouterGroup) {
	r := UserRouter{}

	g.GET("/list", r.UserList)
}

//
// UserList TODO 此处的返回值，根据项目需求返回响应的结构
//  @Description: 查询用户列表
//  @receiver r
//  @param c
//
func (r *UserRouter) UserList(c *gin.Context) {
	var userListReq req.UserList
	if err := c.ShouldBind(userListReq); err != nil {
		c.JSON(http.StatusOK, "")
		return
	}

	resp, err := service.UserService.UserList(userListReq)
	if err != nil {
		c.JSON(http.StatusOK, "")
		return
	}

	c.JSON(http.StatusOK, resp)
}
