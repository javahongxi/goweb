package controllers

import (
	"github.com/astaxie/beego"
	"goweb/models"
)

type UserController struct {
	beego.Controller
}

func Add(c *UserController) error {
	age, _ := c.GetInt16("age")
	user := models.User{
		Name: c.GetString("name"),
		Age:  age,
	}
	return models.Insert(&user)
}

func Delete(c *UserController) error {
	id, _ := c.GetInt("id")
	return models.Delete(id)
}

func Update(c *UserController) error {
	age, _ := c.GetInt16("age")
	user := models.User{
		Name: c.GetString("name"),
		Age:  age,
	}
	return models.Update(&user)
}

func Find(c *UserController) models.User {
	id, _ := c.GetInt("id")
	user, _ := models.Select(id)
	return user
}
