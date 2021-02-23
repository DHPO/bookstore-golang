package dao

import "bookstore/bookstore/model"

type BookDao interface {
	FindOne(id int) (model.Book, error)
	GetBooks() []model.Book
}
