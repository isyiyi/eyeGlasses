package model

import "time"

type Users struct {
	Id         int       `orm:"column(id)"`
	UserName   string    `orm:"column(username)"`
	Password   string    `orm:"column(password)"`
	Gender     string    `orm:"column(gender)"`
	Address    string    `orm:"column(address)"`
	PhoneNum   string    `orm:"column(phone_num)"`
	Birth      time.Time `orm:"column(birth)"`
	CreateTime time.Time `orm:"column(create_time)"`
	IsDelete   int       `orm:"column(is_delete)"`
}
