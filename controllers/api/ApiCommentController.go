/**
接口controller
*/
package api

import (
	"github.com/astaxie/beego/orm"
	"goblog/controllers"
	"goblog/models"
	"strconv"
)

func init() {

}

//评论
type CommentController struct {
	controllers.BaseController
}

//获取内容
func (this *CommentController) Get() {
	var (
		rs models.Json
	)
	id, _ := this.GetInt64("id")
	fid, _ := this.GetInt64("fid")
	uid, _ := this.GetInt64("uid")
	offset, _ := this.GetInt64("offset")
	comment := new(models.Comment)
	user := new(models.User)

	//获取列表
	where := make(map[string]int64)
	order := make(map[string]string)

	where["id"] = id
	where["fid"] = fid
	where["uid"] = uid
	where["offset"] = offset

	order["order"] = this.GetString("order")
	order["by"] = this.GetString("by")
	page, _ := this.GetInt64("page")
	rs = comment.Search(where, order, page)

	//组合用户信息
	rs1 := rs.Data.(map[string]interface{})["list"].([]orm.Params)
	for k, v := range rs1 {
		uid1, _ := strconv.ParseInt(v["uid"].(string), 10, 0)
		userinfo := user.GetUserForUid(uid1)
		rs1[k]["user"] = userinfo.Data
	}

	this.Data["json"] = rs
	this.ServeJson()
}

//添加
func (this *CommentController) Post() {
	comment := new(models.Comment)
	fid, _ := this.GetInt64("fid")
	uid, _ := this.GetInt64("uid")
	context := this.GetString("context")

	user := new(models.User)

	//验证token
	utoken := this.GetString("utoken")
	rs := user.CheckToken(uid, utoken, "user")
	if rs.Status == false {
		this.Data["json"] = rs
		this.ServeJson()
	}

	data := models.Comment{Fid: fid, Uid: uid, Type: "blog", Context: context}
	rs = comment.Add(data)
	this.Data["json"] = rs
	this.ServeJson()
}

//删除
func (this *CommentController) Delete() {
	comment := new(models.Comment)
	user := new(models.User)

	//验证token
	uid, _ := this.GetInt64("uid")
	utoken := this.GetString("utoken")
	rs := user.CheckToken(uid, utoken, "user")
	if rs.Status == false {
		this.Data["json"] = rs
		this.ServeJson()
	}

	id, _ := this.GetInt64("id")
	where := map[string]int64{"id": id}
	rs = comment.Del(where)
	this.Data["json"] = rs
	this.ServeJson()
}
