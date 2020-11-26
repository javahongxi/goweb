package tmp

import "github.com/astaxie/beego/orm"

type User struct {
	Id         int
	Name       string
	Age        int16
	Gender     int8
	Address    string
	HasJob     bool
	CreateDate int64
}

func init() {
	orm.RegisterModel(new(User))
}

func Insert(user *User) error {
	o := orm.NewOrm()
	_, err := o.Insert(&user)
	return err
}

func Delete(id int) error {
	user := User{Id: id}
	o := orm.NewOrm()
	_, err := o.Insert(&user)
	return err
}

func Update(user *User) error {
	o := orm.NewOrm()
	_, err := o.Update(&user)
	return err
}

func Select(id int) (User, error) {
	user := User{Id: id}
	o := orm.NewOrm()
	err := o.Read(&user)
	return user, err
}
