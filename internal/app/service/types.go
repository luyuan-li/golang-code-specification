package service

var (
	UserService IUserService
)

func Init() {
	UserService = NewUserService()
}
