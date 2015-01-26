<{template "header.tpl" .}>

<div class="am-g am-g-fixed blog-g-fixed am-margin-top-xl">
	<h2 class="am-margin-left"><a href="#" onclick="blog_list(1)">管理后台</a></h2>
	<div class="am-u-md-10">
    <span id="list-blog-admin-html"></span>
    <script id="list-blog-admin-template" type="text/x-handlebars-template">
	{{#if indexpage}}
        <table class="am-table am-table-bordered am-table-striped am-table-hover">
            <thead>
                <tr>
                    <th width="">标题</th>
                    <th width="190px">最后更新 <i class="am-icon-sort-down am-text-danger"></i></th>
                    <th width="150px">操作</th>
                </tr>
            </thead>
            <tbody>
			{{#each list}}
            <tr><td><a href="/#blog_show-{{id}}" target="_blank">{{title}}</a></td><td>{{updatetime}}</td><td><a href="#blog_edit-{{id}}" onclick="blog_edit({{id}})">编辑</a>&nbsp;|&nbsp;<a href="javascript:" onclick="blog_del({{id}})">删除</a></td></tr>
			{{/each}}
            </tbody>
        </table>
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

		
		<form class="am-form" id="bolg-form" data-id="{{id}}">
		  <fieldset>
		
			<div class="am-form-group">
			  <label for="doc-ipt-title-1">标题</label>
			  <input type="text" class="title" name="title" value="{{title}}" id="doc-ipt-title-1" placeholder="输入标题">
			</div>
		
			<div class="am-form-group">
			  <label for="doc-ta-1">{{#if id}}编辑{{else}}添加{{/if}}内容</label>
			  <textarea class="context" name="context" rows="20" placeholder="支持makedown语法" id="doc-ta-1">{{context}}</textarea>
			</div>
		
			<p><span class="am-btn am-btn-primary" id="blog-post-btn">提交</span>&nbsp;&nbsp;<a type="submit" class="am-btn am-btn-default" href="javascript:" onclick='if(confirm("确定放弃当前编辑的内容吗?")){ blog_list(1);location.hash="#"}'>取消</a></p>
		  </fieldset>
		</form>

		
	{{/if}}
    </script>
    </div>
    <div class="am-u-md-2">
        <div class="am-panel am-panel-default" data-am-sticky>
          <div class="am-panel-hd">操作面板</div>
          <div class="am-panel-bd">
            	<p><a href="#blog_add" onclick="blog_add()"><i class="am-icon-plus"></i>添加</a></p>
          </div>
        </div>    
    </div>
 

</div>

<link rel="stylesheet" href="/static/css/admin.css"/>
<script src="/static/js/admin.js"></script>

<{template "footer.tpl" .}>