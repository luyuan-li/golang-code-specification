package repository

import "gorm.io/gorm"

const (
	CollectionNameUser = "user"
)

type IUserRepo interface {
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{db}
}

type userRepo struct {
	db *gorm.DB
}
