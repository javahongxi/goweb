package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         uint64    `orm:"column(id)"`
	Username   string    `orm:"column(username)"`
	Nickname   string    `orm:"column(nickname)"`
	Gender     uint8     `orm:"column(gender)"`
	Age        uint16    `orm:"column(age)"`
	CreateDate time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateDate time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func InsertUser(user *User) error {
	o := orm.NewOrm()
	_, err := o.Insert(&user)
	return err
}

func DeleteUser(id uint64) error {
	user := User{Id: id}
	o := orm.NewOrm()
	_, err := o.Delete(&user)
	return err
}

func UpdateUser(user *User) error {
	o := orm.NewOrm()
	_, err := o.Update(user)
	return err
}

func SelectUser(id uint64) (User, error) {
	user := User{Id: id}
	o := orm.NewOrm()
	err := o.Read(&user)
	return user, err
}
