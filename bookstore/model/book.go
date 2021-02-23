package model

type Book struct {
	ID int `gorm:"primarykey" json:"bookId"`
	Isbn string `json:"isbn"`
	Name string `json:"name"`
	BookType string `gorm:"column:type" json:"type"`
	Author string `json:"author"`
	Price float32 `json:"price"`
	Description string `json:"description"`
	Inventory int `json:"inventory"`
	Image string `json:"image"`
}

func (Book) TableName() string {
	return "book"
}
