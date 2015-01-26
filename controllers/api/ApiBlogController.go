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

//博客内容
type BlogContextController struct {
	controllers.BaseController
}

//获取内容
func (this *BlogContextController) Get() {
	var (
		rs models.Json
	)
	id, _ := this.GetInt64(":id")
	blog := new(models.Blog)
	if id > 0 {
		//获取单条内容
		rs = blog.GetInfo(id)
		if rs.Status == true {
			blog.ViewnumAdd(id) //阅读数加1
		}
	} else {
		//获取列表
		where := make(map[string]interface{})

		where["order"] = this.GetString("order")
		where["by"] = this.GetString("by")
		page, _ := this.GetInt64("page")
		rs = blog.Search(where, page)

	}
	this.Data["json"] = rs
	this.ServeJson()
}

//添加
func (this *BlogContextController) Post() {
	user := new(models.User)

	//验证token
	uid, _ := this.GetInt64("uid")
	atoken := this.GetString("atoken")
	rs := user.CheckToken(uid, atoken, "admin")
	if rs.Status == false {
		this.Data["json"] = rs
		this.ServeJson()
	}

	blog := new(models.Blog)
	data := models.Blog{Title: this.GetString("title"), Context: this.GetString("context")}
	rs = blog.Add(data)
	this.Data["json"] = rs
	this.ServeJson()
}

//更新
func (this *BlogContextController) Put() {
	user := new(models.User)

	//验证token
	uid, _ := this.GetInt64("uid")
	atoken := this.GetString("atoken")
	rs := user.CheckToken(uid, atoken, "admin")
	if rs.Status == false {
		this.Data["json"] = rs
		this.ServeJson()
	}

	blog := new(models.Blog)
	id, _ := this.GetInt64("id")
	data := models.Blog{Title: this.GetString("title"), Context: this.GetString("context")}
	rs = blog.Update(id, data)
	this.Data["json"] = rs
	this.ServeJson()
}

//删除
func (this *BlogContextController) Delete() {

	user := new(models.User)

	//验证token
	uid, _ := this.GetInt64("uid")
	atoken := this.GetString("atoken")
	rs := user.CheckToken(uid, atoken, "admin")
	if rs.Status == false {
		this.Data["json"] = rs
		this.ServeJson()
	}

	blog := new(models.Blog)
	id, _ := this.GetInt64(":id")
	rs = blog.Del(id)
	this.Data["json"] = rs
	this.ServeJson()
}
