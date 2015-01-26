package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type Comment struct {
	Id         int64
	Fid        int64
	Uid        int64
	Type       string
	Context    string
	Createtime string
}

func init() {

	orm.RegisterModel(new(Comment))

}

func (this *Comment) Search(where map[string]int64, order map[string]string, page int64) Json {
	var (
		data, data1                       []orm.Params
		rs                                Json
		nextPage, prePage, flag, totalNum int64
		offset                            int64 = 2
		orderby, whereSql                 string
	)

	o := orm.NewOrm()

	if page < 1 {
		page = 1
	}
	if where["offset"] > 0 {
		offset = where["offset"]
	}

	//分页
	limit := offset * (page - 1)

	//排序
	order1 := order["order"]
	by1 := order["by"]
	if strings.EqualFold(order1, "") {
		order1 = "id"
	}
	if strings.EqualFold(by1, "") {
		by1 = "desc"
	}
	orderby = order1 + " " + by1

	//条件
	whereSql = "1"
	if where["uid"] > 0 {
		whereSql = whereSql + " and uid=" + strconv.FormatInt(where["uid"], 10)
	}
	if where["fid"] > 0 {
		whereSql = whereSql + " and fid=" + strconv.FormatInt(where["fid"], 10)
	}
	if where["id"] > 0 {
		whereSql = whereSql + " and id=" + strconv.FormatInt(where["id"], 10)
	}

	sql := "SELECT * FROM comment WHERE " + whereSql
	sqlCount := "SELECT COUNT(*) AS num FROM comment WHERE " + whereSql
	//计算总数
	o.Raw(sqlCount).Values(&data1)
	s := data1[0]["num"].(string)
	totalNum, _ = strconv.ParseInt(s, 10, 64)

	//分页查找
	sql = sql + " order by " + orderby + " limit " + strconv.FormatInt(limit, 10) + "," + strconv.FormatInt(offset, 10)
	o.Raw(sql).Values(&data)
	//
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
func (this *Comment) Add(data Comment) Json {
	rs := Json{Status: true}
	if len(data.Context) < 3 || len(data.Context) > 10000 {
		rs.Status = false
		rs.Message = "评论长度必须是3~10000字内"
		return rs
	}
	if len(data.Type) <= 0 {
		rs.Status = false
		rs.Message = "评论类型不能空"
		return rs
	}
	if data.Fid <= 0 {
		rs.Status = false
		rs.Message = "评论对象ID不能空"
		return rs
	}
	if data.Uid <= 0 {
		rs.Status = false
		rs.Message = "评论UID不能空"
		return rs
	}

	o := orm.NewOrm()
	c := new(Comment)
	c.Fid = data.Fid
	c.Uid = data.Uid
	c.Context = data.Context
	c.Createtime = time.Now().Format("2006-01-02 15:04:05")
	c.Type = data.Type
	id, _ := o.Insert(c)

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

//删除评论
func (this *Comment) Del(where map[string]int64) Json {
	rs := Json{Status: true}
	o := orm.NewOrm()
	data := Comment{}

	if where["id"] > 0 {
		data.Id = where["id"]
	}
	if where["fid"] > 0 {
		data.Id = where["id"]
	}
	if where["uid"] > 0 {
		data.Id = where["id"]
	}
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
