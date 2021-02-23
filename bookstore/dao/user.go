package dao

import "bookstore/bookstore/model"

type UserDao interface {
	CheckUser(username, password string) (model.UserAuth, error)
}
