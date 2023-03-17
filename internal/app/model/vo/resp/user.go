package resp

import "github.com/golang-code-specification/internal/app/enum"

type (
	User struct {
		ID         uint            `json:"id"`
		Name       string          `json:"name"`
		Age        int64           `json:"age"`
		Mobile     string          `json:"mobile"`
		UserStatus enum.UserStatus `json:"user_status"`
	}

	UserList []User
)
