package controllers

//博客管理
type BlogAdminController struct {
	BaseController
}

//博客首页
func (this *BlogAdminController) Get() {
	this.TplNames = "admin/blog.tpl"
}
