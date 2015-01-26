# goblog

##介绍

这是基于beego的go语言个人blog系统,目前项目处于起步阶段!随时可能增减功能!欢迎大家一起来完善! let's go!

对于个人博客sqlite已经可以满足,但为了提供更高的性能程序支持mysql和sqlite3

# 安装(install)

## sqlite3数据库

1.下载源代码,然后把文件解压到$GOPATH/src/goblog

2.配置./conf/app.conf 去掉mysql配置,SqliteDbName为sqlite路径

3.进入 $GOPATH/src/goblog 执行 go install & go run main.go 

(_注意:放到goblog/目录,不是github.com/目录,这个以后会完善_)

4.在浏览器输入 http://localhost:8080

## mysql数据库

1.下载源代码,然后把文件解压到$GOPATH/src/goblog

2.配置./conf/app.conf mysql信息

3.导入./database文件下goblog.sql语句

4.进入 $GOPATH/src/goblog 执行 go install & go run main.go 

(_注意:放到goblog/目录,不是github.com/目录,这个以后会完善_)

5.在浏览器输入 http://localhost:8080

## 用到的第三方程序

1. https://github.com/astaxie/beego
2. https://github.com/go-sql-driver/mysql
3. http://amazeui.org
4. http://jquery.com
5. http://www.handlebarsjs.com
 
##demo

http://goblog.tangfengqiao.cn


