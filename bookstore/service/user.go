package service

import "bookstore/bookstore/model"

type UserService interface {
	CheckUser(username, password string) (model.UserAuth, error)
}
