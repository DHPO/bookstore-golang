package service

import (
	"bookstore/bookstore/dao"
	"bookstore/bookstore/model"
)

type UserServiceImpl struct {
	dao dao.UserDao
}

func (us *UserServiceImpl) CheckUser(username, password string) (model.UserAuth, error) {
	return us.dao.CheckUser(username, password)
}

func NewUserServiceImpl(dao dao.UserDao) UserService {
	return &UserServiceImpl{
		dao: dao,
	}
}
