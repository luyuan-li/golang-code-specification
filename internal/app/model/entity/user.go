package entity

type User struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Name   string `gorm:"column:name;type:varchar(32);comment:姓名" json:"name"`
	Age    int64  `gorm:"column:age;type:int;comment:年龄" json:"age"`
	Mobile string `gorm:"column:mobile;type:varchar(32);comment:手机号" json:"mobile"`
}

func (u User) TableName() string {
	return "user"
}
