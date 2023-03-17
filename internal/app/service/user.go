package service

import (
	"github.com/golang-code-specification/internal/app/model/vo/resp"
	"github.com/golang-code-specification/internal/app/repository"
)

type IUserService interface {
	UserList() (resp.UserListResp, error)
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
//  @return resp.UserListResp
//  @return error
//
func (srv *userService) UserList() (resp.UserListResp, error) {
	users, err := repository.UserRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var userListResp resp.UserListResp
	for _, user := range users {
		var userResp resp.UserResp
		userResp.ID = user.ID
		userResp.Name = user.Name
		userResp.Age = user.Age
		userResp.Mobile = user.Mobile
		userListResp = append(userListResp, userResp)
	}

	return userListResp, nil
}
