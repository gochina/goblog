$(document).ready(function(e) {

	atoken = cookie.get("atoken");
	uid = cookie.get("uid");
	if(!atoken){
		window.location.href="/#blog_adminlogin";
		return false;
	}
				 
	Gtitle = "管理后台"
	//加载url
	var hash = window.location.hash;
	var arr = hash.split('-');
	var m = arr[0].replace("#","");
	if(m){
		var a = parseInt(arr[1]);
		eval(m+"("+a+")");	
	}else{
		blog_list(1);
	}
	
	

	document.title = Gtitle
	
	//保存内容
	$(this).on('click','#blog-post-btn',function(){
		var obj = $('#bolg-form');
		var title = obj.find('.title').val();
		var context = obj.find('.context').val();
		var id = obj.data('id');
		var url;
		if(!title){
			alert('请输入标题');
			return false;	
		}
		if(!context){
			alert('请输入内容');
			return false;	
		}
		url = "/api/blog/";
		if(id){
			//编辑
			type = "PUT"
		}else{
			//新增
			type = "POST"
		}
		url = "/api/blog/";
		$.ajax({
			url:url,
			data:{uid:uid,atoken:atoken,title:title,context:context,id:id},
			type:type,
			success:function(rs){
				alert(rs.Message);
				if(rs.Status == true){
					$.AMUI.store.clear();//清除浏览器缓存
					document.location.hash = "#"
					blog_list(1);	
				}
			}
		});			
	});
});


//博客列表
function blog_list(page){
	var url ="/api/blog?page="+page
	
	$.get(url,{uid:uid,atoken:atoken},function(rs){
		var obj = $("#list-blog-admin-template");
		var template = Handlebars.compile(obj.html());
		if(rs.Status){
			rs.Data.indexpage = 1; //首页
			var html = template(rs.Data);
		}else{
			html = rs.Message;	
		}
		document.title = (rs.Data.page?"第"+rs.Data.page+"页":"")+"-博客列表-" + Gtitle
		$("#list-blog-admin-html").html(html);		
	});	
}

//编辑博客内容
function blog_edit(id){
	var url ="/api/blog/"+id
	$.get(url,function(rs){
		var obj = $("#list-blog-admin-template");
		var template = Handlebars.compile(obj.html());
		if(rs.Status){
			rs.Data.indexpage = 0; //不是首页
			var html = template(rs.Data);
		}else{
			html = rs.Message;
		}
		$("#list-blog-admin-html").html(html);		
	});	
}
//添加博客内容
function blog_add(){

	var obj = $("#list-blog-admin-template");
	var template = Handlebars.compile(obj.html());
	var html = template({});
	$("#list-blog-admin-html").html(html);		

}

//删除博客内容
function blog_del(id){
	var url ="/api/blog/"+id+"?uid="+uid+"&atoken="+atoken;
	if(!confirm("确定删除吗")){
		return false;
	}
	$.ajax({
		url:url,
		type:"DELETE",
		success:function(rs){
			if(rs.Status){
				$.AMUI.store.clear();//清除浏览器缓存
				blog_list(1);
			}else{
				alert(rs.Message);
			}
		}
	});			
}