package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "goblog/routers"
)

func init() {
	MysqlUser := beego.AppConfig.String("MysqlUser")
	MysqlPassword := beego.AppConfig.String("MysqlPassword")
	MysqlHost := beego.AppConfig.String("MysqlHost")
	MysqlPort := beego.AppConfig.String("MysqlPort")
	MysqlDbName := beego.AppConfig.String("MysqlDbName")
	if len(MysqlHost) > 0 {
		var dns string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", MysqlUser, MysqlPassword, MysqlHost, MysqlPort, MysqlDbName)
		orm.RegisterDataBase("default", "mysql", dns)
	} else {
		SqliteDbname := beego.AppConfig.String("SqliteDbname")
		orm.RegisterDataBase("default", "sqlite3", SqliteDbname)
	}

}
func main() {
	beego.Run()
}
