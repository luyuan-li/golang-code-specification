package service

import (
	"github.com/golang-code-specification/internal/app/model/vo/req"
	"github.com/golang-code-specification/internal/app/model/vo/resp"
	"github.com/golang-code-specification/internal/app/repository"
)

type IUserService interface {
	UserList(userListReq req.UserList) (resp.UserList, error)
}

type userService struct {
}

func NewUserService() IUserService {
	return &userService{}
}

//
// UserList
//  @Description: 查询用户列表
//  @receiver srv
//  @param userListReq
//  @return resp.UserList
//  @return error
//
func (srv *userService) UserList(userListReq req.UserList) (resp.UserList, error) {
	users, err := repository.UserRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var userListResp resp.UserList
	for _, user := range users {
		var userResp resp.User
		userResp.ID = user.ID
		userResp.Name = user.Name
		userResp.Age = user.Age
		userResp.Mobile = user.Mobile
		userResp.UserStatus = user.UserStatus
		userListResp = append(userListResp, userResp)
	}

	return userListResp, nil
}
