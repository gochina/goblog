<{template "header.tpl" .}>

<header class="am-topbar" data-am-sticky>
  <div class="am-topbar-brand" onclick="blog_list(1)">
    <a href="#" id="logo-html"></a>
  </div>

  <button class="am-topbar-btn am-topbar-toggle am-btn am-btn-sm am-btn-success am-show-sm-only" data-am-collapse="{target: '#topbar-collapse'}"><span class="am-sr-only">导航切换</span> <span class="am-icon-bars"></span></button>

  <div class="am-collapse am-topbar-collapse" id="topbar-collapse">
    <span id="top-user-html"></span>
    <script id="top-user-template" type="text/x-handlebars-template">
	{{#if uid}}
		<ul class="am-nav am-nav-pills am-topbar-nav am-topbar-right admin-header-list">
		
		  <li id="top-user-menu">
			<a id="right-menu-link" href="javascript:;">
			  <span class="am-icon-user"></span> {{username}}
			</a> 
		  </li>
		</ul>
	
	{{else}}
	 
		<div class="am-topbar-right">
			<button class="am-btn am-btn-secondary am-topbar-btn am-btn-xs am-radius register-link"><span class="am-icon-pencil"></span>&nbsp;注册</button>
		</div>
		<div class="am-topbar-right">
			<button class="am-btn am-btn-primary am-topbar-btn am-btn-xs am-radius login-link"><span class="am-icon-user"></span>&nbsp;登录</button>
		</div>
	
	{{/if}}
	</script>        
  </div>
</header>

<div class="am-g am-g-fixed blog-g-fixed am-margin-top-xl">
  <div class="am-u-md-8">
	<span id="list-blog-html">
	<{if .rs.IsList}> 
		<{range $v := .rs.Data}>
		<article class="blog-main">
		  <h3 class="am-article-title blog-title">
			<a href="/blog/<{$v.id}>"><{$v.title}></a>
		  </h3>
	
		  <div class=" blog-content">
		  <{substr $v.intro 0 50}>
		  </div>
		  <div class="ccc am-margin-top am-text-xs"><{$v.createtime}> - 来自:博客</div>
		</article>
		<hr class="am-article-divider blog-hr"> 
		<{end}>
	<{else}>
	<h1><{.rs.Data.title}></h1>
	<article class="am-padding-top" id="blog-context" >
	<{.rs.Data.context}>
    </article>
	<{end}>
 
	</span>  
	<script id="list-blog-template" type="text/x-handlebars-template">
	{{#if indexpage}}
		{{#each list}}
		<article class="blog-main">
		  <h3 class="am-article-title blog-title">
			<a href="#blog_show-{{id}}" onclick="blog_show({{id}})">{{title}}</a>
		  </h3>
	
		  <div class=" blog-content">
		  {{title}}
		  </div>
		  <div class="ccc am-margin-top am-text-xs">{{createtime}} - 来自:博客</div>
		</article>
		<hr class="am-article-divider blog-hr"> 
		{{/each}}
		
		<ul class="am-pagination blog-pagination">
		{{#compare page 1}}
			<li class="am-pagination-prev"><a href="#" onclick="blog_list(1)">首页</a></li>
		{{/compare}}
		{{#if prepage}}
		  <li class="am-pagination-prev"><a href="#blog_list-{{prepage}}" onclick="blog_list({{prepage}})">«  上一页</a></li>
		{{/if}}
		{{#if nextpage}}
		  <li class="am-pagination-prev"><a href="#blog_list-{{nextpage}}" class=" am-text-danger" onclick="blog_list({{nextpage}})">下一页  »</a></li>
		{{/if}}
		</ul>
	{{else}}
		<div class="am-cf am-article">
			<h1 class="am-text-center am-margin-xl">{{title}}</h1>
			<div class="ccc am-margin-top am-margin-botton am-text-xs">阅读:{{viewnum}} 发表于:{{createtime}}</div>
			<hr /> 
		</div>
		<article class="am-padding-top" id="blog-context" data-blogid="{{id}}">
		{{{context}}}
	    </article>	
		<div class="am-margin-top am-margin-bottom">
			<a href="#" class="am-btn am-btn-secondary" onclick="blog_list(1)"><i class="am-icon-arrow-circle-left"></i> 返回</a>
		</div>	
	
	{{/if}}
	</script>
    <span id="list-blogcomment-html"></span>
	<script type="text/x-handlebars-template" id="list-blogcomment-template">
		<div class="am-margin-top">
			<legend>评论</legend>
			<ul class="am-comments-list">
			  {{#each list}}
			  <li class="am-comment">
				  <a href="">
					<img class="am-comment-avatar" alt="{{user.username}}" src="{{#if user.avatar}}{{user.avatar}}{{else}}/static/img/avatar.gif{{/if}}" />
				  </a>
				
				  <div class="am-comment-main">
					<header class="am-comment-hd">
					  <div class="am-comment-meta">
					  	<div class="am-fl"> 
						<a href="#link-to-user" class="am-comment-author">{{user.username}}</a> 
						评论于 <time datetime="">{{createtime}}</time>
						</div>
						{{#me uid ../UID}}
						<a href="javascript:" class="am-fr blog-comment-del-btn" data-cid="{{id}}" data-fid="{{fid}}">删除</a>
						{{/me}}
					  </div>
					</header>
				
					<div class="am-comment-bd">{{context}}</div> <!-- 评论内容 -->
				  </div>
			  </li>
			  {{/each}}
			  		
				<ul class="am-pagination blog-pagination" style="text-align:center">
				{{#if nextpage}}
				  <li class="am-pagination-next"><a href="javascript:" class=" am-text-danger" onclick="blogcomment_list({{fid}},{{nextpage}})">下一页  »</a></li>
				{{/if}}
				{{#if prepage}}
				  <li class="am-pagination-next am-margin-right"><a href="javascript:" onclick="blogcomment_list({{fid}},{{prepage}})">«  上一页</a></li>
				{{/if}}
				{{#compare page 1}}
					<li class="am-pagination-next am-margin-right"><a href="javascript:" onclick="blogcomment_list({{fid}},1)">首页</a></li>
				{{/compare}}
				</ul>

			  <div class="am-comment">
				
				  <div class="am-comment-main">
					<header class="am-comment-hd">
					  <div class="am-comment-meta"> 
						发表评论
					  </div>
					</header>
				
					<div class="am-comment-bd">
						<textarea  class="" id="blog-comment-context" style="width:100% ;border:0px;" name="context" rows="5" placeholder="在这输入您的评论" ></textarea>
						<div class="am-margin-top am-margin-bottom">
							<a href="javascript:" class="am-btn am-fr am-btn-secondary am-radius"  id="blog-comment-btn"><i class="am-icon-arrow-circle-up"></i> 提交评论</a>
						</div>
						
					</div> <!-- 评论内容 -->
				  </div>
			  </div>

			</ul>
			
		</div>
	</script>
 
  </div>

  <div class="am-u-md-4 blog-sidebar">
    <div class="am-panel-group">
      <section class="am-panel am-panel-default">
        <div class="am-panel-hd">关于我</div>
        <div class="am-panel-bd" id="abuot-html">
        
        </div>
      </section>
	<span id="hot-blog-html"></span>  
	<script id="hot-blog-template" type="text/x-handlebars-template">
      {{#if list}}
      <div class="am-margin"></div>
      <section class="am-panel am-panel-default">
        <div class="am-panel-hd">热门排行</div>
        <ul class="am-list blog-list">
		{{#each list}}
          <li><a href="#blog_show-{{id}}" onclick="blog_show({{id}})">{{title}}</a></li>
		{{/each}}
        </ul>
      </section>
	  {{/if}}
	</script>
    </div>
  </div>

</div>

<{template "blog/model.tpl" .}>
<link rel="stylesheet" href="/static/css/blog.css"/>
<script src="/static/markdown/showdown.js"></script>
<script src="/static/js/blog.js"></script>

<{template "footer.tpl" .}>