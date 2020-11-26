package tmp

import (
	"github.com/astaxie/beego"
	"goweb/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Add() {
	age, _ := c.GetInt16("age")
	user := models.User{
		Name: c.GetString("name"),
		Age:  age,
	}
	err := models.Insert(&user)
	if err != nil {
		c.Ctx.WriteString("op success")
	} else {
		c.Ctx.WriteString("op failed")
	}
}

func (c *UserController) Delete() {
	id, _ := c.GetInt("id")
	err := models.Delete(id)
	if err != nil {
		c.Ctx.WriteString("op success")
	} else {
		c.Ctx.WriteString("op failed")
	}
}

func (c *UserController) Update() {
	age, _ := c.GetInt16("age")
	user := models.User{
		Name: c.GetString("name"),
		Age:  age,
	}
	err := models.Update(&user)
	if err != nil {
		c.Ctx.WriteString("op success")
	} else {
		c.Ctx.WriteString("op failed")
	}
}

func (c *UserController) Find() {
	id, _ := c.GetInt("id")
	user, err := models.Select(id)
	if err != nil {
		c.Ctx.WriteString("op success")
	} else {
		c.Ctx.WriteString(user.Name)
	}
}
