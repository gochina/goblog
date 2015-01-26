package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goblog/models"
)

//博客
type BlogController struct {
	BaseController
}

//博客首页
func (this *BlogController) Get() {
	type rs struct {
		IsList bool
		Data   interface{}
		Status bool
		Title  string
	}
	var result rs

	blog := new(models.Blog)
	id, _ := this.GetInt64(":id")
	if id > 0 {
		//获取单条内容
		rs := blog.GetInfo(id)
		if rs.Status == true {
			result.Data = rs.Data
			result.Title = rs.Data.(orm.Params)["title"].(string) + "-" + beego.AppConfig.String("appname")
		} else {
			result.Status = false
		}

		result.IsList = false
		this.Data["rs"] = result

	} else {
		//获取列表
		where := make(map[string]interface{})
		where["order"] = "id"
		where["by"] = "desc"
		page, _ := this.GetInt64("page")
		rs := blog.Search(where, page)
		result.Title = "第1页-博客列表-" + beego.AppConfig.String("appname")
		result.IsList = true
		if rs.Status == true {
			result.Data = rs.Data.(map[string]interface{})["list"].([]orm.Params)
			result.Status = true
		} else {
			result.Status = false
		}
		this.Data["rs"] = result

	}
	this.TplNames = "blog/index.tpl"
}
