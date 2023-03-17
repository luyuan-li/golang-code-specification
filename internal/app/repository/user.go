package repository

import (
	"github.com/golang-code-specification/internal/app/model/entity"
	"gorm.io/gorm"
)

type IUserRepo interface {
	FindAll() ([]entity.User, error)
	FindByName() (entity.User, error)
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{db}
}

type userRepo struct {
	db *gorm.DB
}

//
// FindAll
//  @Description: 查询所有用户
//  @receiver repo
//  @return []entity.User
//  @return error
//
func (repo *userRepo) FindAll() ([]entity.User, error) {
	var res []entity.User
	err := repo.db.Find(&res).Error
	return res, err
}

//
// FindByName
//  @Description: 根据名称查询用户
//  @receiver repo
//  @return entity.User
//  @return error
//
func (repo *userRepo) FindByName() (entity.User, error) {
	return entity.User{}, nil
}
