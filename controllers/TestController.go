package controllers

import (
	"goblog/models"
)

func init() {

}

//登录
type TestController struct {
	BaseController
}

func (this *TestController) Get() {

	user := new(models.User)
	rs := user.Login(this.GetString("u"), this.GetString("p"))
	this.Data["json"] = rs
	this.ServeJson()
}
