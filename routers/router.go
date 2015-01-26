// @APIVersion 1.0.0
// @Title goblog 主程序
// @Description 模板输出
// @Contact www.tangfengqiao.cn

package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers"
)

func init() {
	beego.Router("/admin", &controllers.BlogAdminController{})
	beego.Router("/", &controllers.BlogController{})
	beego.Router("/blog/:id", &controllers.BlogController{})
	beego.Router("/upload", &controllers.UploadController{}, "post:Upload")
	beego.Router("/upload/avatar", &controllers.UploadController{}, "*:Avatar")
}
