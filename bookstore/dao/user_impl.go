package dao

import (
	"bookstore/bookstore/model"
	"gorm.io/gorm"
)

type UserDaoImpl struct {
	db *gorm.DB
}

func NewUserDaoImpl(db *gorm.DB) UserDao {
	return &UserDaoImpl{
		db: db,
	}
}

func (this *UserDaoImpl) CheckUser(username, password string) (model.UserAuth, error) {
	auth := model.UserAuth{}
	err := this.db.Where(&model.UserAuth{
		Username: username,
		Password: password,
	}).First(&auth).Error

	return auth, err
}
