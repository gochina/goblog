# goblog
## 安装 (install)

1.获取源代码 go get github.com/gochina/goblog

2.配置./conf/app.conf mysql信息

3.导入./database文件下goblog.sql语句

4.进入 $GOPATH/src/goblog 执行 go install & go run main.go 
(_注意:放到goblog/目录,不是github.com/目录,这个以后会完善_)

5.在浏览器输入 http://localhost:8080

## 第三方包

1. go get github.com/astaxie/beego
2. go get github.com/go-sql-driver/mysql

