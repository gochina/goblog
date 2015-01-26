package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Json struct {
	Status  bool
	Message interface{}
	Data    interface{}
	Error   error
	Atoken  string
}

type User struct {
	Uid        int64 `orm:"pk"`
	Username   string
	Password   string
	Email      string
	Phone      string
	Qq         string
	Createtime string
	Updatetime string
	Sex        string
	Avatar     string
}

func init() {
	orm.RegisterModel(new(User))
	orm.Debug = true
}

//登录
func (this *User) Login(username string, password string) Json {
	o := orm.NewOrm()
	var (
		rs   Json = Json{Status: true}
		data []orm.Params
	)

	num, _ := o.Raw("SELECT * FROM user WHERE (username = ? or email=?) and password=?", username, username, this.MakePassword(password)).Values(&data)
	if num > 0 {
		rs.Message = "登录成功"
		tmp := make(map[string]interface{})
		tmp["uid"], _ = strconv.ParseInt(data[0]["uid"].(string), 10, 0)
		tmp["username"] = username
		rs.Data = tmp
	} else {
		rs.Status = false
		rs.Message = "登录失败!请检查账号密码是否正确"
		rs.Data = ""
	}

	return rs
}

//注册
func (this *User) Register(data User) Json {
	var (
		rs   Json = Json{Status: true}
		user User
	)
	o := orm.NewOrm()
	//检查用户名格式
	rs = this.CheckUsernameType(data.Username)
	if rs.Status == false {
		return rs
	}

	//判断用户名是否存在
	rs = this.HasUsername(data.Username)
	if rs.Status == true {
		rs.Status = false
		return rs
	}

	//检查邮箱格式
	rs = this.CheckEmailType(data.Email)
	if rs.Status == false {
		return rs
	}

	//判断邮箱是否存在
	rs = this.HasEmail(data.Email)
	if rs.Status == true {
		rs.Status = false
		return rs
	}

	//检查密码格式
	rs = this.CheckPasswordType(data.Password)
	if rs.Status == false {
		return rs
	}

	//如果设置手机号,检查格式
	if len(data.Phone) > 0 {
		rs := this.CheckPhoneType(data.Phone)
		if rs.Status == false {
			return rs
		}
	}

	//如果设置qq,检查格式
	if len(data.Qq) > 0 {
		rs := this.CheckQqType(data.Qq)
		if rs.Status == false {
			return rs
		}
	}

	user.Username = data.Username
	user.Password = this.MakePassword(data.Password)
	user.Email = data.Email
	user.Phone = data.Phone
	user.Qq = data.Qq
	user.Sex = data.Sex
	user.Uid = 0
	user.Createtime = time.Now().Format("2006-01-02 15:04:05")
	id, err := o.Insert(&user)
	if nil != err {
		rs.Status = false
		rs.Message = "用户注册失败"
		rs.Error = err
		rs.Data = 0
	} else {
		rs.Message = "用户注册成功"
		tmp := make(map[string]interface{})
		tmp["uid"] = id
		tmp["username"] = data.Username
		rs.Data = tmp
	}

	return rs

}

// 按用户名查找用户
func (this *User) GetUserForName(key string) Json {
	o := orm.NewOrm()
	var (
		data User
		rs   Json
	)

	data = User{Username: key}
	o.Read(&data, "username")
	rs.Data = data
	return rs
}

// 按uid查找用户
func (this *User) GetUserForUid(uid int64) Json {
	var (
		maps []orm.Params
		rs   Json
	)
	o := orm.NewOrm()

	o.Raw("SELECT uid,username,sex,email,phone,qq,createtime,updatetime,avatar FROM user where uid=?", uid).Values(&maps)
	if len(maps) > 0 {
		rs.Status = true
		rs.Data = maps[0]
	} else {
		rs.Status = false
	}
	return rs
}

// 修改用户信息
func (this *User) UpdateInfo(uid int64, data map[string]string) Json {
	var rs Json = Json{Status: true}

	if len(data["phone"]) > 0 {
		rs := this.CheckPhoneType(data["phone"])
		if rs.Status == false {
			return rs
		}
	}

	if len(data["qq"]) > 0 {
		rs := this.CheckQqType(data["qq"])
		if rs.Status == false {
			return rs
		}
	}

	o := orm.NewOrm()
	data2 := orm.Params{}
	for k, v := range data {
		if len(v) >= 0 {
			data2[k] = v
		}
	}
	_, err := o.QueryTable("user").Filter("uid", uid).Update(data2)
	if err != nil {
		rs.Status = false
		rs.Message = "用户信息修改失败"
	} else {
		rs.Message = "用户信息修改成功"
	}
	return rs
}

// 修改头像
func (this *User) UpdateAvatar(uid int64, avatar string) Json {
	var rs Json = Json{Status: true}

	o := orm.NewOrm()
	data2 := orm.Params{}

	data2["avatar"] = avatar
	_, err := o.QueryTable("user").Filter("uid", uid).Update(data2)
	if err != nil {
		rs.Status = false
		rs.Message = "头像修改失败"
	} else {
		rs.Message = "头像修改成功"
	}
	return rs
}

// 修改邮箱
func (this *User) UpdateEmail(uid int64, password string, email string) Json {
	var rs Json = Json{Status: true}
	o := orm.NewOrm()

	//检查邮箱格式
	rs = this.CheckEmailType(email)
	if rs.Status == false {
		return rs
	}

	//判断邮箱是否存在
	rs = this.HasEmail(email)
	if rs.Status == true {
		rs.Status = false
		return rs
	}

	user := User{Uid: uid}
	err := o.Read(&user, "uid")
	if nil != err {
		rs.Status = false
		rs.Message = "用户不存在"
		rs.Error = err
	} else {

		if strings.EqualFold(user.Password, this.MakePassword(password)) == true {

			num, err := o.QueryTable("user").Filter("uid", uid).Update(orm.Params{"email": email})

			if nil != err {
				rs.Status = false
				rs.Message = "更新失败"
				rs.Error = err
			} else if 0 == num {
				rs.Status = false
				rs.Message = "没更新"
			} else {
				rs.Message = "邮箱更新成功"
			}
		} else {
			rs.Status = false
			rs.Message = "密码错误"
		}
	}
	return rs
}

// 设置密码
func (this *User) SetPassword(uid int64, password string) Json {
	var rs Json = Json{Status: true}
	o := orm.NewOrm()

	num, err := o.QueryTable("user").Filter("id", uid).Update(orm.Params{
		"password": password,
	})

	if 0 == num {
		rs.Status = false
		rs.Message = "没有操作"
	} else if nil != err {
		rs.Status = false
		rs.Message = "密码设置失败"
		rs.Error = err
	} else {
		rs.Message = "密码设置成功"
	}
	return rs
}

// 修改密码
func (this *User) UpdatePassword(uid int64, oldPassword string, newPassword string) Json {
	var rs Json = Json{Status: true}
	o := orm.NewOrm()

	rs = this.CheckPasswordType(newPassword)
	if rs.Status == false {
		return rs
	}

	user := User{Uid: uid}
	err := o.Read(&user, "uid")
	if nil != err {
		rs.Status = false
		rs.Message = "用户不存在"
		rs.Error = err
	} else {

		if strings.EqualFold(user.Password, this.MakePassword(oldPassword)) == true {
			num, err := o.QueryTable("user").Filter("uid", uid).Update(orm.Params{"password": this.MakePassword(newPassword)})
			if 0 == num {
				rs.Status = false
				rs.Message = "没有操作"
			} else if nil != err {
				rs.Status = false
				rs.Message = "密码修改失败"
				rs.Error = err
			} else {
				rs.Message = "密码修改成功"
			}
		} else {
			rs.Status = false
			rs.Message = "旧密码错误"
		}
	}
	return rs
}

//判断用户名是否存在
func (this *User) HasUsername(username string) Json {
	var (
		rs Json = Json{Status: true}
	)
	o := orm.NewOrm()
	num, _ := o.QueryTable("user").Filter("username", username).Count()
	if num > 0 {
		rs.Status = true
		rs.Data = num
		rs.Message = "此用户名已存在"
	} else {
		rs.Status = false
		rs.Data = 0
		rs.Message = "此用户名不存在"
	}
	return rs
}

//判断邮箱是否存在
func (this *User) HasEmail(email string) Json {
	var (
		rs Json = Json{Status: true}
	)
	o := orm.NewOrm()
	num, _ := o.QueryTable("user").Filter("email", email).Count()
	if num > 0 {
		rs.Status = true
		rs.Data = num
		rs.Message = "此用邮箱已存在"
	} else {
		rs.Status = false
		rs.Data = 0
		rs.Message = "此邮箱不存在"
	}
	return rs
}

//检查密码格式
func (this *User) CheckPasswordType(password string) Json {
	var rs Json = Json{Status: true}
	reg := regexp.MustCompile(`^[a-zA-Z0-9_\.@]{6,50}$`)
	result := reg.MatchString(password)
	if !result {
		rs.Status = false
		rs.Message = password + `密码必须是由长度6~50的字母数字_.@组成`
	}
	return rs
}

//检查邮箱格式
func (this *User) CheckEmailType(email string) Json {
	var rs Json = Json{Status: true}
	reg := regexp.MustCompile(`^[a-zA-Z0-9_\.]+@[a-zA-Z0-9-]+\.[a-zA-Z]+$`)
	result := reg.MatchString(email)
	if !result {
		rs.Status = false
		rs.Message = "邮箱格式错误！"
	}
	return rs
}

//检查用户名格式
func (this *User) CheckUsernameType(username string) Json {
	var rs Json = Json{Status: true}
	reg := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{4,49}$`)
	result := reg.MatchString(username)
	if !result {
		rs.Status = false
		rs.Message = "用户名格式错误！用户名必须以字母开头，长度在5~50之间，只能包含字母、数字和下划线"
	}
	return rs
}

//验证token
func (this *User) CheckToken(uid int64, token string, key string) Json {
	var rs Json = Json{Status: true}
	str := this.MakeToken(uid, key)
	if strings.EqualFold(str, token) {
		rs.Status = true
		rs.Message = "token验证通过"
	} else {
		rs.Status = false
		rs.Message = "token验证失败,请重新验证"
	}
	return rs
}

//检查手机格式
func (this *User) CheckPhoneType(phone string) Json {
	var rs Json = Json{Status: true}

	reg := regexp.MustCompile(`^1[3|4|5|8]\d{9}$`)
	result := reg.MatchString(string(phone))
	if !result {
		rs.Status = false
		rs.Message = "手机号码格式错误"
	} else {
		rs.Message = "手机号码格式正确"
	}
	return rs
}

//检查qq格式
func (this *User) CheckQqType(qq string) Json {
	var rs Json = Json{Status: true}

	reg := regexp.MustCompile(`^[0-9]{3,12}$`)
	result := reg.MatchString(string(qq))
	if !result {
		rs.Status = false
		rs.Message = "qq号码格式错误"
	} else {
		rs.Message = "qq号码格式正确"
	}
	return rs
}

//生成密码
func (this *User) MakePassword(password string) string {
	password = password + "~!@#$%^&*()_+"
	p := md5.New()
	p.Write([]byte(password))

	return hex.EncodeToString(p.Sum(nil))
}

//生成用户token
func (this *User) MakeToken(uid int64, key string) string {
	var (
		str string
	)
	str = key + strconv.FormatInt(uid, 10)
	p := md5.New()
	p.Write([]byte(str))

	return hex.EncodeToString(p.Sum(nil))
}
