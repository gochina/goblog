/**
接口controller
*/
package api

import (
	"goblog/controllers"
	"goblog/models"
)

func init() {
}

//会话
type SessionController struct {
	controllers.BaseController
}

//创建会话
func (this *SessionController) Post() {
	data := make(map[string]interface{})
	user := new(models.User)
	username := this.GetString("username")
	password := this.GetString("password")
	rs := user.Login(username, password)
	if rs.Status == true {
		uid := rs.Data.(map[string]interface{})["uid"].(int64)
		data["username"] = rs.Data.(map[string]interface{})["username"]
		data["uid"] = uid
		data["utoken"] = user.MakeToken(uid, "user")
		rs.Data = data

		//判断是否管理员
		if controllers.AdminUID == uid {
			rs.Atoken = user.MakeToken(uid, "admin")
		}
	}
	this.Data["json"] = rs
	this.ServeJson()

}

//删除会话
func (this *SessionController) Delete() {
	var (
		rs models.Json
	)
	this.DelSession("uid")
	rs.Status = true
	rs.Message = "退出成功"
	this.Data["json"] = rs
	this.ServeJson()

}

//用户信息
type UserinfoController struct {
	controllers.BaseController
}

//获取用户信息
func (this *UserinfoController) Get() {
	var (
		uid int64
	)
	uid, _ = this.GetInt64("uid")

	user := new(models.User)
	rs := user.GetUserForUid(uid)

	this.Data["json"] = rs
	this.ServeJson()
}

//更新用户信息
func (this *UserinfoController) Put() {

	do := this.GetString("do")
	uid, _ := this.GetInt64("uid")
	utoken := this.GetString("utoken")

	user := new(models.User)

	//验证token
	rs := user.CheckToken(uid, utoken, "user")
	if rs.Status == false {
		this.Data["json"] = rs
		this.ServeJson()
	}

	switch do {
	case "password":
		//编辑密码
		rs = user.UpdatePassword(uid, this.GetString("oldpassword"), this.GetString("newpassword"))
	case "email":
		//编辑邮箱
		rs = user.UpdateEmail(uid, this.GetString("password"), this.GetString("email"))
	case "avatar":
		rs = user.UpdateAvatar(uid, this.GetString("avatar"))
	default:
		//默认编辑用户信息
		data := make(map[string]string)
		data["sex"] = this.GetString("sex")
		data["phone"] = this.GetString("phone")
		data["qq"] = this.GetString("qq")

		rs = user.UpdateInfo(uid, data)
	}
	this.Data["json"] = rs
	this.ServeJson()

}

//注册
func (this *UserinfoController) Post() {

	user := new(models.User)

	data := models.User{
		Username: this.GetString("username"),
		Password: this.GetString("password"),
		Email:    this.GetString("email"),
		Sex:      this.GetString("sex"),
		Phone:    this.GetString("phone"),
		Qq:       this.GetString("qq"),
	}
	rs := user.Register(data)
	if rs.Status == true {
		data := make(map[string]interface{})
		uid := rs.Data.(map[string]interface{})["uid"].(int64)
		data["uid"] = uid
		data["username"] = rs.Data.(map[string]interface{})["username"]
		data["utoken"] = user.MakeToken(uid, "user")
		rs.Data = data
	}
	this.Data["json"] = rs
	this.ServeJson()

}
