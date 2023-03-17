package entity

import "github.com/golang-code-specification/internal/app/enum"

type User struct {
	ID         uint            `gorm:"primarykey" json:"id"`
	Name       string          `gorm:"column:name;type:varchar(32);comment:姓名" json:"name"`
	Age        int64           `gorm:"column:age;type:int;comment:年龄" json:"age"`
	Mobile     string          `gorm:"column:mobile;type:varchar(32);comment:手机号" json:"mobile"`
	UserStatus enum.UserStatus `gorm:"column:user_status;type:int;comment:手机号" json:"user_status"`
}

func (u User) TableName() string {
	return "user"
}
