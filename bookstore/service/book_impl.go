package service

import (
	"bookstore/bookstore/dao"
	"bookstore/bookstore/model"
)

type BookServiceImpl struct {
	dao dao.BookDao
}

func NewBookServiceImpl(dao dao.BookDao) BookService {
	return &BookServiceImpl{
		dao: dao,
	}
}

func (bs *BookServiceImpl) FindBookById(id int) (model.Book, error) {
	return bs.dao.FindOne(id)
}

func (bs *BookServiceImpl) GetBooks() []model.Book {
	return bs.dao.GetBooks()
}
