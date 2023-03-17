package enum

type (
	UserStatus    int
	ServiceStatus int
)

const (
	UserStatusEnable   UserStatus = 1 //启用
	UserStatusDisabled UserStatus = 2 //禁用

	ServiceStatusEnable   ServiceStatus = 1 //启用
	ServiceStatusDisabled ServiceStatus = 2 //禁用
)

//TODO 可以将枚举写成map，如下
var UserStatusMap = map[UserStatus]string{
	UserStatusEnable:   "启用",
	UserStatusDisabled: "禁用",
}
