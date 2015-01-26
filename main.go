package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "goblog/routers"
)

func init() {
	var dns string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", beego.AppConfig.String("mysqluser"), beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlhost"), beego.AppConfig.String("mysqlport"), beego.AppConfig.String("mysqldb"))
	orm.RegisterDataBase("default", "mysql", dns)

}
func main() {
	beego.Run()
}
