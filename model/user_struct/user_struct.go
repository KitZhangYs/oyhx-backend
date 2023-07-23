package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"primarykey" gorm:"size:30"`
	Password string `json:"password" gorm:"size:30"`
	MailAddr string `json:"mail_addr" gorm:"primarykey" gorm:"size:30"`
}

type UserCode struct {
	gorm.Model
	Username string `json:"username" gorm:"primarykey" gorm:"size:30"`
	Code     int    `json:"code" gorm:"primarykey"`
}
