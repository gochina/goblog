$(document).ready(function(e) {

	Gtitle = " goblog"
		//加载url
	var hash = window.location.hash;
	var arr = hash.split('-');
	var arr1 = hash.split('_');
	var m = arr[0].replace("#", "");
	if (m && arr1[0] == "#blog") {
		var a = parseInt(arr[1]);
		eval(m + "(" + a + ")");
	} else {
		blog_list(1);
	}

	//获取热门博客列表
	blog_hotlist(1);

	checkUserBar();



	//注册
	$('.register-btn').on('click', function() {
		user.register();
	});

	//提交评论
	$(this).on('click', '#blog-comment-btn', function() {
		var fid = $("#blog-context").data('blogid');
		var uid = cookie.get("uid");
		var utoken = cookie.get("utoken");
		var context = $('#blog-comment-context').val();

		var url = "/api/comment";

		//判断是否登录
		if (!uid) {
			alert("请登录后再评论");
			return false;
		}

		//评论内容
		if (context.length < 3 || context.length > 10000) {
			alert("评论长度必须是3~10000字内");
			return false;
		}

		//判断fid
		if (!fid) {
			alert("评论对象ID不能空");
			return false;
		}
		$.post(url, {
			utoken: utoken,
			uid: uid,
			fid: fid,
			context: context
		}, function(rs) {
			if (rs.Status == false) {
				alert(rs.Message);
			} else {
				closeAll();
				alert(rs.Message);
				$('#blog-comment-context').val('');
				blogcomment_list(fid, 1, true);
			}
		});
	});

	//删除博客评论
	$(this).on('click', '.blog-comment-del-btn', function() {
		var cid = $(this).data('cid');
		var fid = $(this).data('fid');
		var uid = cookie.get("uid");
		var admincp = $(this).data('admincp');

		var utoken = cookie.get("utoken");
		var url = "/api/comment?id=" + cid + "&uid=" + uid + "&utoken=" + utoken;
		if (!uid) {
			alert("请登录");
			return false;
		}
		if (!cid) {
			alert("评论对象不能空");
			return false;
		}
		if (!fid) {
			alert("博客ID不能空");
			return false;
		}
		if (!confirm("确定删除此评论吗?")) {
			return false;
		}
		$.ajax({
			url: url,
			type: "DELETE",
			success: function(rs) {
				if (rs.Status == true) {
					if (admincp) {
						mycomment_list(1);
					} else {
						blogcomment_list(fid, 1, true);
					}

				} else {
					alert(rs.Message);
				}
			}
		});
	});

	//提交登录
	$(this).on('click', '.login-btn', function() {
		user.login();
	});

	//退出
	$(this).on("click", ".logout-btn", function() {
		user.logout();
	});

	//编辑链接
	$(this).on('click', '.edit-link', function() {
		editUserModel();
	});

	//编辑链接
	$(this).on('click', '.comment-link', function() {
		myCommentModel();
	});

	//注册链接
	$(this).on('click', '.register-link', function() {
		$('#login-model').find(".am-close").click();
		$('#register-model').modal({
			closeViaDimmer: false
		});
	});

	//登录链接
	$(this).on('click', '.login-link', function() {
		$('#register-model').find(".am-close").click();
		$('#login-model').modal({
			closeViaDimmer: false
		});
	});

	//编辑用户基本信息
	$('.edit-user-btn').on('click', function() {
		user.updateInfo();
	});

	//修改密码
	$('.edit-password-btn').on('click', function() {
		user.updatePwd();
	});

	//修改email
	$('.edit-email-btn').on('click', function() {
		user.updateEmail();
	});


	//打开侧栏
	$(this).on('click', '#right-menu-link', function() {
			getRightMenuData();
			$(".am-offcanvas").offCanvas('open')
		})
		//基本信息
	$('#logo-html').html("Go语言博客");
	$("#abuot-html").html("<p>GoBlog是一款go语言开发的个人博客系统，程序基于beego框架开发。</p><p>我们崇尚简约不简单！抛开复杂的功能，发现博客简单之美。</p>");

	//上传头像
	$(this).on('click', '#avatar', function() {
		var uid = cookie.get("uid");
		$('#uploadUid').val(uid)
		$('.upload-file').click();
	});
	$('.upload-file').on('change', function() {
		var editForm = $('.upload-form');
		$.AMUI.progress.start();
		$('.upload-form').submit();
	});
});


//上传图片成功回调
function uploadCallBack(path, status, message) {
	$.AMUI.progress.done();
	if (status == "false") {
		alert(message)
		return false;
	}
	var path = path;
	var url = "/api/user?do=avatar";
	if (path.indexOf("://") > 0) {
		path += "!avatar"
	}
	$.ajax({
		url: url,
		data: {
			avatar: path
		},
		type: "PUT",
		success: function(rs) {
			if (rs.Status == true) {
				$('#avatar').find('img').attr('src', path);
			}
			alert(rs.Message)
		}
	});

}

//关闭所有窗口/层
function closeAll() {
	$(".am-popup").modal("close");
	$(".am-offcanvas").offCanvas('close');
}

function checkLogin() {
	var uid = cookie.get("uid");
	if (!uid > 0) {
		checkUserBar();
		closeAll();
		alert("请登录");
		return false
	}

}

//判断登录栏:op==open 打开右侧菜单栏
function checkUserBar(op) {
	var key;
	var obj = $("#top-user-template");
	var template = Handlebars.compile(obj.html());
	var uid = cookie.get("uid");
	var utoken = cookie.get("utoken");
	var username = cookie.get("username");
	var data = {};

	if (uid && utoken) {
		if (op == "open") {
			setTimeout(function() {
				$('#right-menu-link').click();
			}, 300);
		}
		data.uid = uid;
		data.username = username;
	}
	$("#top-user-html").html(template(data));
}

//加载右边菜单数据
function getRightMenuData(nocache) {

	var obj = $("#ucenter-template");
	var template = Handlebars.compile(obj.html());
	var uid = cookie.get("uid");
	var key = "cache_user_" + uid;
	var url = "/api/user?uid=" + uid;
	var atoken = cookie.get("atoken");

	rs = store.get(key) ? store.get(key) : {};
	if (rs.Status == true && !nocache) {
		if (!rs.Data.avatar) {
			rs.Data.avatar = "/static/img/avatar.gif"
		}
		rs.atoken = atoken;
		$('#ucenter-html').html(template(rs));
	} else {
		$.get(url, function(rs) {
			rs.atoken = atoken;
			store.set(key, rs);
			if (rs.Status == true) {
				if (!rs.Data.avatar) {
					rs.Data.avatar = "/static/img/avatar.gif"
				}
			}
			$('#ucenter-html').html(template(rs));
		});
	}
}


//编辑用户信息窗口
function editUserModel() {
	$(".am-offcanvas").offCanvas('close')
	var obj = $('#edit-model');
	obj.modal({
		closeViaDimmer: false
	});
	var url = "/api/user";
	$.get(url, function(rs) {
		if (!rs.Status) {
			return false;
		}
		obj.find("[name=phone]").val(rs.Data.phone);
		obj.find("[name=qq]").val(rs.Data.qq);
		obj.find("[name=sex][value=" + rs.Data.sex + "]").attr("checked", "checked");
	});
}

//我的评论面板
function myCommentModel() {
	$(".am-offcanvas").offCanvas('close');
	var obj = $('#comment-model').modal({
		closeViaDimmer: 0
	});
	mycomment_list(1);
}

//加载我的评论数据
function mycomment_list(page) {
	if (!page > 0) {
		page = 1;
	}
	var uid = cookie.get("uid");
	var url = "/api/comment?offset=5&uid=" + uid + "&page=" + page
	var obj = $("#list-mycomment-template");
	var template = Handlebars.compile(obj.html());

	$.get(url, function(rs) {
		rs.Data.UID = uid;
		var html = template(rs.Data);
		$("#list-mycomment-html").html(html);
	});
}

//博客列表
function blog_list(page) {
	if (!page > 0) {
		page = 1;
	}
	var url = "/api/blog?page=" + page
	var obj = $("#list-blog-template");
	var template = Handlebars.compile(obj.html());
	var key = "cache_bloglist_" + page;
	var rs = store.get(key) ? store.get(key) : {};

	//只读首页缓存
	if (rs.Status == true && page == 1) {
		if (rs.Status) {
			rs.Data.indexpage = 1; //首页
			var html = template(rs.Data);
		} else {
			html = rs.Message;
		}
		document.title = (rs.Data.page ? "第" + rs.Data.page + "页" : "") + "-博客列表-" + Gtitle
		$("#list-blog-html").html(html);
	} else {
		$.get(url, function(rs) {
			if (rs.Status) {
				if (rs.Status == true && page == 1) {
					store.set(key, rs, 300)
				}
				rs.Data.indexpage = 1; //首页
				var html = template(rs.Data);
			} else {
				html = rs.Message;
			}
			document.title = (rs.Data.page ? "第" + rs.Data.page + "页" : "") + "-博客列表-" + Gtitle
			$("#list-blog-html").html(html);
		});
	}
	$("#list-blogcomment-html").html(""); //清空评论
}

//博客评论列表
function blogcomment_list(fid, page, nocache) {
	if (!page > 0) {
		page = 1;
	}

	var url = "/api/comment?fid=" + fid + "&page=" + page
	var obj = $("#list-blogcomment-template");
	var template = Handlebars.compile(obj.html());
	var key = "cache_blogcommentlist_fid" + fid + "_page" + page;
	var rs = store.get(key) ? store.get(key) : {};
	var uid = cookie.get("uid");

	//只读首页缓存
	if (rs.Status == true && page == 1 && !nocache) {

		rs.Data.UID = uid;
		var html = template(rs.Data);

		$("#list-blogcomment-html").html(html);
	} else {
		$.get(url, function(rs) {
			if (rs.Status) {
				//只缓存首页
				if (rs.Status == true && page == 1) {
					store.set(key, rs)
				}
			} else {
				html = rs.Message;
			}
			rs.Data.UID = uid;
			var html = template(rs.Data);
			$("#list-blogcomment-html").html(html);
		});
	}
}

//推荐博客列表
function blog_hotlist() {
	var url = "/api/blog?order=viewnum&by=desc&offset=10"
	var obj = $("#hot-blog-template");
	var template = Handlebars.compile(obj.html());
	var key = "cache_bloghot";

	rs = store.get(key) ? store.get(key) : {};
	if (rs.Status == true) {
		if (rs.Status) {
			var html = template(rs.Data);
		} else {
			html = rs.Message;
		}
		$("#hot-blog-html").html(html);
	} else {
		$.get(url, function(rs) {
			if (rs.Status) {
				store.set(key, rs);
				var html = template(rs.Data);
			} else {
				html = rs.Message;
			}
			$("#hot-blog-html").html(html);
		});
	}
}


//博客内容
function blog_show(id) {
	var url = "/api/blog/" + id
	var obj = $("#list-blog-template");
	var template = Handlebars.compile(obj.html());
	var key = "cache_blog_" + id;

	$(window).scrollTop(0);

	rs = store.get(key) ? store.get(key) : {};
	if (rs.Status == true) {
		if (rs.Status == true) {
			rs.Data.indexpage = 0; //不是首页
			converter = new Showdown.converter();
			rs.Data.context = converter.makeHtml(rs.Data.context);
			var html = template(rs.Data);
			document.title = rs.Data.title + "-" + Gtitle
		} else {
			html = rs.Message;
		}
		$("#list-blog-html").html(html);
	} else {
		$.get(url, function(rs) {
			if (rs.Status == true) {
				store.set(key, rs);
				rs.Data.indexpage = 0; //不是首页
				converter = new Showdown.converter();
				rs.Data.context = converter.makeHtml(rs.Data.context);
				var html = template(rs.Data);
				document.title = rs.Data.title + "-" + Gtitle
			} else {
				html = rs.Message;
			}

			$("#list-blog-html").html(html);
		});
	}
	blogcomment_list(id, 1);
}

function blog_adminlogin() {
	alert("请用管理员账号登录");
	blog_list(1);
	window.location.hash = ""
}