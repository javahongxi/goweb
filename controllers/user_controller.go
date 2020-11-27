package controllers

import (
	"github.com/astaxie/beego"
	"goweb/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Add() {
	gender, _ := c.GetUint8("gender")
	age, _ := c.GetUint16("age")
	user := models.User{
		Username: c.GetString("username"),
		Nickname: c.GetString("nickname"),
		Gender:   gender,
		Age:      age,
	}
	err := models.InsertUser(&user)
	if err == nil {
		c.Ctx.WriteString("op success")
	} else {
		c.Ctx.WriteString("op failed")
	}
}

func (c *UserController) Delete() {
	id, _ := c.GetUint64("id")
	err := models.DeleteUser(id)
	if err == nil {
		c.Ctx.WriteString("op success")
	} else {
		c.Ctx.WriteString("op failed")
	}
}

func (c *UserController) Update() {
	id, _ := c.GetUint64("id")
	user := models.User{Id: id}
	if username := c.GetString("username"); username != "" {
		user.Username = username
	}
	if nickname := c.GetString("nickname"); nickname != "" {
		user.Nickname = nickname
	}

	err := models.UpdateUser(&user)
	if err == nil {
		c.Ctx.WriteString("op success")
	} else {
		c.Ctx.WriteString("op failed")
	}
}

func (c *UserController) Find() {
	id, _ := c.GetUint64("id")
	user, err := models.SelectUser(id)
	if err == nil {
		c.Ctx.WriteString(user.Nickname)
	} else {
		c.Ctx.WriteString("no user")
	}
}
