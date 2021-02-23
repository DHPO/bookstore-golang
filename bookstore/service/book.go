package service

import "bookstore/bookstore/model"

type BookService interface {
	FindBookById(id int) (model.Book, error)
	GetBooks() []model.Book
}
