package model

import "encoding/gob"

type UserAuth struct {
	ID int `gorm:"primarykey;column:user_id"`
	Username string `json:"username"`
	Password string `json:"-"`
	UserType int `gorm:"column:user_type" json:"userType"`
}

func (UserAuth) TableName() string {
	return "user_auth"
}

func init() {
	gob.Register(UserAuth{})
}