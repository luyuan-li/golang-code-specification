package repository

import "gorm.io/gorm"

var (
	UserRepo IUserRepo
)

func Init(db *gorm.DB) {
	UserRepo = NewUserRepo(db)
}
