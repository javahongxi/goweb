package controllers

import (
	"github.com/astaxie/beego"
	result "goweb/common"
	"goweb/models"
)

type UserController struct {
	beego.Controller
}

type userParam struct {
	Id       uint64 `form:"id"`
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Gender   uint8  `form:"gender"`
	Age      uint16 `form:"age"`
}

func (c *UserController) Hello() {
	c.Ctx.WriteString("Hello, " + c.GetString("name", "Kitty"))
}

func (c *UserController) Add() {
	//gender, _ := c.GetUint8("gender")
	//age, _ := c.GetUint16("age")

	var r result.Result
	userParam := userParam{}
	var err error
	if err = c.ParseForm(&userParam); err == nil {
		user := models.User{
			Username: userParam.Username,
			Nickname: userParam.Nickname,
			Gender:   userParam.Gender,
			Age:      userParam.Age,
		}
		err = models.InsertUser(&user)
	}

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
