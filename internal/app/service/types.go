package service

var (
	UserService IUserService
)

func InitService() {
	UserService = NewUserService()
}
