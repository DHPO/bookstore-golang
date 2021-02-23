package dao

import (
	"bookstore/bookstore/model"
	"gorm.io/gorm"
)

type BookDaoImpl struct {
	db *gorm.DB
}


func NewBookDaoImpl(db *gorm.DB) BookDao {
	return &BookDaoImpl{
		db: db,
	}
}

func (this *BookDaoImpl) FindOne(id int) (model.Book, error) {
	book := model.Book{}
	err := this.db.First(&book, id).Error
	return book, err
}


func (this *BookDaoImpl) GetBooks() []model.Book {
	books := []model.Book{}
	this.db.Find(&books)
	return books
}
