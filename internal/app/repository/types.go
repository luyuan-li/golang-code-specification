package repository

import "gorm.io/gorm"

var (
	UserRepo IUserRepo
)

func InitRepo(db *gorm.DB) {
	UserRepo = NewUserRepo(db)
}
