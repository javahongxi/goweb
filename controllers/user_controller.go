package controllers

import (
	"github.com/astaxie/beego"
	result "goweb/common"
	"goweb/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Hello() {
	c.Ctx.WriteString("Hello, " + c.GetString("name", "Kitty"))
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

	var r result.Result
	if err == nil {
		r = result.Result{}
	} else {
		r = result.Error(500, err.Error())
	}
	c.Data["json"] = &r
	c.ServeJSON()
}

func (c *UserController) Delete() {
	id, _ := c.GetUint64("id")
	err := models.DeleteUser(id)

	var r result.Result
	if err == nil {
		r = result.Result{}
	} else {
		r = result.Error(500, err.Error())
	}
	c.Data["json"] = &r
	c.ServeJSON()
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

	var r result.Result
	if err == nil {
		r = result.Result{}
	} else {
		r = result.Error(500, err.Error())
	}
	c.Data["json"] = &r
	c.ServeJSON()
}

func (c *UserController) Find() {
	id, _ := c.GetUint64("id")
	user, err := models.SelectUser(id)

	var r result.Result
	if err == nil {
		r = result.Success(&user)
	} else {
		r = result.Error(500, err.Error())
	}
	c.Data["json"] = &r
	c.ServeJSON()
}
