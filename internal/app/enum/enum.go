package enum

type (
	UserStatus int
)

const (
	// EnableUserStatus 启用
	EnableUserStatus UserStatus = 0
	// DisabledUserStatus 禁用
	DisabledUserStatus UserStatus = 1
)
