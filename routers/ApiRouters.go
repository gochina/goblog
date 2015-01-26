// @APIVersion 1.0.0
// @Title goblog API
// @Description go博客api
// @Contact www.tangfengqiao.cn
package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers/api"
)

func init() {
	beego.Router("/api/blog/?:id", &api.BlogContextController{})

	beego.Router("/api/session", &api.SessionController{})
	beego.Router("/api/user", &api.UserinfoController{})

	beego.Router("/api/comment", &api.CommentController{})
}
