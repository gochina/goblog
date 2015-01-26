package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Blog struct {
	Id         int64
	Title      string
	Context    string
	Createtime string
	Updatetime string
	Viewnum    int64
	Intro      string
}

func init() {

	orm.RegisterModel(new(Blog))

}

func (this *Blog) Search(where map[string]interface{}, page int64) Json {
	var (
		data                    []orm.Params
		rs                      Json
		nextPage, prePage, flag int64
		offset                  int64 = 20
		orderby                 string
	)
	if page < 1 {
		page = 1
	}
	limit := offset * (page - 1)

	o := orm.NewOrm()

	//排序
	order := where["order"].(string)
	by := where["by"].(string)
	if strings.EqualFold(order, "") {
		order = "updatetime"
	}
	if strings.EqualFold(by, "") {
		by = "desc"
	}
	orderby = order + " " + by

	sql := "SELECT * FROM blog "
	//计算总数
	totalNum, _ := o.QueryTable("blog").Count()

	//分页查找
	sql = sql + "order by " + orderby + " limit ?,?"
	o.Raw(sql, limit, offset).Values(&data)

	//计算分页
	if 0 == (totalNum % offset) {
		flag = 0
	} else {
		flag = 1
	}
	totalPage := totalNum/offset + flag

	if totalPage <= page {
		nextPage = 0
	} else {
		nextPage = page + 1
	}
	if page > 1 {
		prePage = page - 1
	}
	if totalNum > 0 {
		rs.Status = true
	} else {
		rs.Message = "暂无数据"
		rs.Status = false
		rs.Data = ""
	}
	rs1 := make(map[string]interface{})
	rs1["nextpage"] = nextPage
	rs1["prepage"] = prePage
	rs1["page"] = page
	rs1["list"] = data
	rs1["totalpage"] = totalPage
	rs1["offset"] = offset
	rs.Data = rs1
	return rs
}

//添加
func (this *Blog) Add(data Blog) Json {
	//检查标题格式
	rs := this.CheckTitleType(data.Title)
	if rs.Status == false {
		return rs
	}
	o := orm.NewOrm()
	blog := new(Blog)
	blog.Title = data.Title
	blog.Context = data.Context
	time := time.Now().Format("2006-01-02 15:04:05")
	blog.Createtime = time
	blog.Updatetime = time
	blog.Viewnum = 1
	id, _ := o.Insert(blog)

	if id > 0 {
		rs.Status = true
		rs.Data = id
		rs.Message = "添加成功"
	} else {
		rs.Status = false
		rs.Message = "添加失败"
	}
	return rs
}

// 修改
func (this *Blog) Update(id int64, data Blog) Json {

	//检查标题格式
	rs := this.CheckTitleType(data.Title)
	if rs.Status == false {
		return rs
	}

	rs = this.GetInfo(id)
	if rs.Status == false {
		return rs
	}
	o := orm.NewOrm()
	time := time.Now().Format("2006-01-02 15:04:05")
	_, err := o.QueryTable("blog").Filter("id", id).Update(orm.Params{
		"title":      data.Title,
		"context":    data.Context,
		"updatetime": time,
	})
	if err != nil {
		rs.Status = false
		rs.Message = "修改失败"
	} else {
		rs.Message = "修改成功"
	}
	return rs
}

//获取内容
// 按uid查找用户
func (this *Blog) GetInfo(id int64) Json {
	var (
		data []orm.Params
		rs   Json
	)
	o := orm.NewOrm()

	o.Raw("SELECT * FROM blog where id=?", id).Values(&data)
	if len(data) > 0 {
		rs.Status = true
		rs.Data = data[0]
	} else {
		rs.Status = false
		rs.Message = "内容不存在"
	}
	return rs
}

//检查标题是否合法
func (this *Blog) CheckTitleType(title string) Json {
	rs := Json{Status: true}
	if len(title) <= 0 {
		rs.Status = false
		rs.Message = "标题不能空"
	} else {
		rs.Status = true
		rs.Message = "标题合法"
	}
	return rs
}

//删除
func (this *Blog) Del(id int64) Json {
	rs := Json{Status: true}
	o := orm.NewOrm()
	data := Blog{Id: id}
	num, _ := o.Delete(&data)
	if num > 0 {
		rs.Status = true
		rs.Message = "删除成功"
	} else {
		rs.Status = false
		rs.Message = "删除失败"
	}
	return rs
}

//访问量增加
func (this *Blog) ViewnumAdd(id int64) Json {
	rs := Json{Status: true}
	o := orm.NewOrm()
	orm.Debug = true
	_, err := o.Raw("UPDATE blog SET viewnum=`viewnum`+1 WHERE id=?", id).Exec()
	if err == nil {
		rs.Status = true
		rs.Message = "阅读数增加1"
	} else {
		rs.Status = false
		rs.Message = "阅读数增加失败"
		rs.Error = err
	}
	return rs
}
