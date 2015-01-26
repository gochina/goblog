package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type BaseController struct {
	beego.Controller
}

type MainController struct {
	BaseController
}

type Site struct {
	Title       string
	Keyword     string
	Description string
}

var (
	UID, AdminUID int64
)

func init() {
	AdminUID, _ = strconv.ParseInt(beego.AppConfig.String("AdminUID"), 10, 0)

}
func (this *BaseController) Prepare() {

	uid := this.GetSession("uid")
	if uid != nil {
		UID = uid.(int64)
	} else {
		UID = 0
	}
	fmt.Println(uid)
}

func (this *MainController) Get() {

	this.TplNames = "index.tpl"
}
