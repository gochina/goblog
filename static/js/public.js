$(document).ready(function(e) {


	//aja状态条
	$(this).ajaxStart(function() {
		$.AMUI.progress.start();
	});

	$(this).ajaxComplete(function() {
		$.AMUI.progress.done();
	});



	//注册模板函数
	Handlebars.registerHelper("compare", function(v1, v2, options) {
		v1 = parseInt(v1);
		if (v1 > v2) {
			return options.fn(this);
		} else {
			return options.inverse(this);
		}
	});
	Handlebars.registerHelper("me", function(uid, UID, options) {

		if (uid == UID) {
			return options.fn(this);
		} else {
			return options.inverse(this);
		}
	});

});

//自动删除缓存数据
astroe = $.AMUI.store;
cookie = $.AMUI.utils.cookie;

//cooke过期时间
overday = new Date();
overday.setTime(overday.getTime() + 86400000 * 30); //30天

var store = {
	set: function(key, val, exp) {
		exp = exp ? exp : 3600 * 5;
		var nowtime = new Date().getTime() / 1000;
		astroe.set(key, {
			val: val,
			exp: exp,
			time: nowtime
		});
	},

	get: function(key) {
		var info = astroe.get(key)
		if (!info) {
			return null;
		}
		var nowtime = new Date().getTime() / 1000;
		if (nowtime - info.time > info.exp) {
			return null;
		}

		return info.val
	}
}

var user = {
	logout: function() {
		cookie.unset("uid");
		cookie.unset("utoken");
		cookie.unset("username");
		cookie.unset("atoken");
		checkUserBar();
		closeAll();
		$('.blog-comment-del-btn').remove(); //清除删除评论按钮
		astroe.clear();
		$("[type=password]").val("");
	},
	login: function() {
		var obj = $('#login-form')
		var username = obj.find("[name=username]").val();
		var password = obj.find("[name=password]").val();
		var url = "/api/session";
		$.post(url, {
			username: username,
			password: password
		}, function(rs) {
			if (rs.Status == false) {
				alert(rs.Message);
			} else {
				cookie.set("uid", rs.Data.uid, overday);
				cookie.set("utoken", rs.Data.utoken, overday);
				cookie.set("username", rs.Data.username, overday);
				cookie.set("atoken", rs.Atoken, overday);
				closeAll();
				checkUserBar('open');
			}
		});
	},
	register: function() {
		var obj = $('#register-form')
		var username = obj.find("[name=username]").val();
		var password = obj.find("[name=password]").val();
		var password1 = obj.find("[name=password1]").val();
		var email = obj.find("[name=email]").val();
		var phone = obj.find("[name=phone]").val();
		var sex = obj.find("[name=sex]").val();
		var qq = obj.find("[name=qq]").val();
		var filter;
		var url = "/api/user";

		//判断用户名
		filter = /^[a-zA-Z][a-zA-Z0-9_]{4,49}$/;
		if (!filter.test(username)) {
			alert("用户名格式错误！用户名必须以字母开头，长度在5~50之间，只能包含字母、数字和下划线");
			return false;
		}

		//判断邮箱
		filter = /^[a-zA-Z0-9_\.]+@[a-zA-Z0-9-]+\.[a-zA-Z]+$/;
		if (!filter.test(email)) {
			alert("邮箱格式错误");
			return false;
		}

		//判断密码
		filter = /^[a-zA-Z0-9_\.@]{6,50}$/;
		if (!filter.test(password)) {
			alert("密码格式错误!密码必须是由长度6~50的字母数字_.@组成");
			return false;
		}
		if (password != password1) {
			alert("两次密码不一致");
			return false;
		}
		$.post(url, {
			username: username,
			password: password,
			email: email,
			phone: phone,
			sex: sex,
			qq: qq
		}, function(rs) {
			if (rs.Status == false) {
				alert(rs.Message)
			} else {
				closeAll();
				cookie.set("uid", rs.Data.uid, overday);
				cookie.set("utoken", rs.Data.utoken, overday);
				cookie.set("username", rs.Data.username, overday);
				closeAll();
				checkUserBar('open');
			}
		});
	},
	updatePwd: function() {
		var obj = $('#edit-form')
		var oldpassword = obj.find("[name=oldpassword]").val();
		var newpassword = obj.find("[name=newpassword]").val();
		var repassword = obj.find("[name=repassword]").val();
		var uid = cookie.get("uid");
		var utoken = cookie.get("utoken");
		if (!uid) {
			alert("请登录");
			return false;
		}

		//判断密码
		filter = /^[a-zA-Z0-9_\.@]{6,50}$/;
		if (!filter.test(newpassword)) {
			alert("密码格式错误!密码必须是由长度6~50的字母数字_.@组成");
			return false;
		}
		if (repassword != newpassword) {
			alert("两次密码不一致");
			return false;
		}

		var url = "/api/user?do=password";
		$.ajax({
			url: url,
			data: {
				uid: uid,
				utoken: utoken,
				oldpassword: oldpassword,
				newpassword: newpassword
			},
			type: "PUT",
			success: function(rs) {
				if (rs.Status == true) {
					$("[type=password]").val("");
				}
				alert(rs.Message);
			}
		});

	},
	updateEmail: function() {
		var obj = $('#edit-form')
		var password = obj.find("[name=password]").val();
		var email = obj.find("[name=email]").val();
		var uid = cookie.get("uid");
		var utoken = cookie.get("utoken");
		if (!uid) {
			alert("请登录");
			return false;
		}

		//判断邮箱
		filter = /^[a-zA-Z0-9_\.]+@[a-zA-Z0-9-]+\.[a-zA-Z]+$/;
		if (!filter.test(email)) {
			alert("邮箱格式错误");
			return false;
		}

		var url = "/api/user?do=email";
		$.ajax({
			url: url,
			data: {
				uid: uid,
				utoken: utoken,
				password: password,
				email: email
			},
			type: "PUT",
			success: function(rs) {

				if (rs.Status == true) {
					getRightMenuData(true);
					$("[type=password]").val("");
				}
				alert(rs.Message);
			}
		});
	},
	updateInfo: function() {
		var obj = $('#edit-form')
		var qq = obj.find("[name=qq]").val();
		var phone = obj.find("[name=phone]").val();
		var sex = obj.find("[name=sex]:checked").val();
		var uid = cookie.get("uid");
		var utoken = cookie.get("utoken");
		var url = "/api/user?uid=" + uid + "&utoken=" + utoken;
		$.ajax({
			url: url,
			data: {
				qq: qq,
				phone: phone,
				sex: sex
			},
			type: "PUT",
			success: function(rs) {

				if (rs.Status == true) {
					getRightMenuData(true)
				} else {
					checkLogin();
				}
				alert(rs.Message)
			}
		});
	},
	clearCache: function(uid) {
		var key = "cache_user_" + uid;
		$.AMUI.store.set(key, null);
	}
}