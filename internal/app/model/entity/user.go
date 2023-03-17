package entity

import (
	"github.com/golang-code-specification/internal/app/enum"
)

type User struct {
	ID         uint            `gorm:"primarykey" json:"id"`
	Name       string          `gorm:"column:name;comment:姓名" json:"name"`
	Age        int64           `gorm:"column:age;comment:年龄" json:"age"`
	Mobile     string          `gorm:"column:mobile;comment:手机号" json:"mobile"`
	UserStatus enum.UserStatus `gorm:"column:user_status;comment:手机号" json:"user_status"`
	CreatedAt  int64           `gorm:"column:created_at;comment:创建时间" json:"created_at"`
	UpdatedAt  int64           `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
}

// TableName return table name 'user'
func (u User) TableName() string {
	return "user"
}

// AllColumns return list columns name for table 'user'
func (*User) AllColumns() []string {
	return []string{"id", "name", "age", "mobile", "user_status", "created_at", "updated_at"}
}

// table 'user' columns list struct
type userTableColumns struct {
	ID         string
	Name       string
	Age        string
	Mobile     string
	UserStatus string
	CreatedAt  string
	UpdatedAt  string
}

// table 'user' columns list info
var UserTableColumns = userTableColumns{
	ID:         "user.id",
	Name:       "user.name",
	Age:        "user.age",
	Mobile:     "user.mobile",
	UserStatus: "user.user_status",
	CreatedAt:  "user.created_at",
	UpdatedAt:  "user.updated_at",
}
