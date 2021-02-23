package model

type User struct {
	ID int `gorm:"primarykey;column:user_id"`
	Nickname string `json:"nickname"`
	Name string `json:"name"`
	Tel string `json:"tel"`
	Address string `json:"address"`
}

func (User) TableName() string {
	return "user"
}
