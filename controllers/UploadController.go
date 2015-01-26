package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
	"goblog/plugins/upyun"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

//上传
type UploadController struct {
	BaseController
}

//上传头像
func (this *UploadController) Avatar() {
	var (
		filepath string
		rs       models.Json
	)
	uid, _ := this.GetInt64("uid")
	if uid <= 0 {
		rs.Status = false
		rs.Message = "请传入uid"
		rs.Data = "0"
	} else {
		_, handler, _ := this.GetFile("upload")
		ext := path.Ext(handler.Filename)

		filepath = beego.AppConfig.String("UploadPath") + "avatar/"
		os.MkdirAll(filepath, 0777)
		filepath = filepath + strconv.FormatInt(uid, 10) + ext
		this.SaveToFile("upload", filepath)

		//上传到又拍云
		yunpath, _ := this.UploadUpyun(filepath)
		if len(yunpath) > 2 {
			rs.Data = yunpath
		} else {
			rs.Data = "/" + filepath
		}

		rs.Status = true
		rs.Message = "上传成功"

	}

	this.Data["rs"] = rs
	this.TplNames = "upload/avatar.tpl"
}

//上传图片
func (this *UploadController) Upload() {
	var (
		rs models.Json
	)

	arr := make(map[string]interface{})

	_, handler, _ := this.GetFile("upload")
	filename := strconv.FormatInt(time.Now().Unix(), 10) + path.Ext(handler.Filename)
	filepath := beego.AppConfig.String("UploadPath") + time.Now().Format("2006/01") + "/" //文件上传路径 /uploads/2014/01/
	os.MkdirAll(filepath, 0777)                                                           //创建目录
	this.SaveToFile("upload", filepath+filename)                                          //上传

	//上传到又拍云
	yunpath, err := this.UploadUpyun(filepath + filename)
	arr["upyunPath"] = yunpath
	arr["upyunError"] = err

	arr["path"] = filepath + filename
	rs.Status = true
	rs.Message = "上传成功"
	rs.Data = arr

	this.Data["json"] = rs
	this.ServeJson()
}

//上传到又拍云,路径;文件名
func (this *UploadController) UploadUpyun(filepath string) (string, error) {
	var (
		imgpath string = ""
		err     error
	)

	bucket := beego.AppConfig.String("UpyunBucket")
	username := beego.AppConfig.String("UpyunUsername")
	password := beego.AppConfig.String("UpyunPasswd")
	uploadpath := beego.AppConfig.String("UploadPath")
	host := beego.AppConfig.String("UpyunHost")

	localfile := filepath
	yunfile := strings.Replace(filepath, uploadpath, "", -1)

	//如果设置了又拍云账号密码则上传
	if len(bucket) > 0 && len(username) > 0 && len(password) > 0 && len(host) > 0 {
		u := upyun.NewUpYun(bucket, username, password)
		fi, err := os.Open(localfile)
		if err != nil {
			return "", err
		}
		u.Delete(yunfile)                      //删除线上同名文件
		_, err = u.Put(yunfile, fi, false, "") //上传到云
		if err != nil {
			return "", err
		}
		imgpath = host + yunfile
	}
	return imgpath, err
}
