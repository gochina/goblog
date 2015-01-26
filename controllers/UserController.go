package controllers

import (
	"goblog/models"
)

func init() {

}

//登录
type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	userInfo := new(models.User).GetUserForUid(UID)
	this.Data["user"] = &userInfo
	this.Data["title"] = "login"
	this.TplNames = "user/login.tpl"
}

//注册
type RegisterController struct {
	BaseController
}

func (this *RegisterController) Get() {
	this.Data["title"] = "login"
	this.TplNames = "user/register.tpl"
}
