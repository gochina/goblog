<div class="am-popup" id="login-model">
  <div class="am-popup-inner">
    <div class="am-popup-hd">
      <h4 class="am-popup-title">用户登录</h4>
      <span data-am-modal-close
            class="am-close">&times;</span>
    </div>
    <div class="am-popup-bd am-padding">
    <form method="post" class="am-form am-padding-xl" id="login-form">
      <div class="form-login-username">
          <label class="am-form-label" for="username">账号:<span class="msg"></span></label>
          <input type="text" name="username" class="am-form-field" placeholder="用户名、邮箱" />
      </div>
      <div class="form-login-password">
          <label class="am-form-label" for="password">密码:<span class="msg"></span></label>
          <input type="password" name="password" class="am-form-field" placeholder="密码" />
      </div>
      <br />
      <div class="am-cf">
        <span class="am-btn am-btn-primary am-btn-sm am-fl login-btn">登 录 <i class="am-icon-circle-o-notch am-icon-spin" style="display:none;"></i></span>
        <span class="am-btn am-btn-sm am-fr register-link">注 册</span>
      </div>
    </form>         
    </div>
    
  </div>
</div>
<!-- end login -->

<div class="am-popup" id="register-model">
  <div class="am-popup-inner">
    <div class="am-popup-hd">
      <h4 class="am-popup-title">用户注册</h4>
      <span data-am-modal-close class="am-close">&times;</span>
    </div>
    <div class="am-popup-bd am-padding-top">
    <form method="post" class="am-form am-padding-top" id="register-form">
      <div class="form-register-username">
          <label class="am-form-label" for="username">账号:<span class="msg"></span></label>
          <input type="text" name="username" class="am-form-field" placeholder="6~50字符" />
      </div>
      <div class="form-register-email">
          <label class="am-form-label" for="email">邮箱:<span class="msg"></span></label>
          <input type="text" name="email" class="am-form-field" placeholder="建议QQ邮箱" />
      </div>
      <div class="form-register-password">
          <label class="am-form-label" for="password">密码:<span class="msg"></span></label>
          <input type="password" name="password" class="am-form-field" placeholder="6~50个字符" />
      </div>
      <div class="form-register-password1">
          <label class="am-form-label" for="password1">重复:<span class="msg"></span></label>
          <input type="password" name="password1" class="am-form-field" placeholder="重复密码" />
      </div>
       <div class="form-register-qq">
          <label class="am-form-label" for="qq">QQ:<span class="msg"></span></label>
          <input type="text" name="qq" class="am-form-field" placeholder="纯数字" />
      </div>
      <div class="form-register-phone">
          <label class="am-form-label" for="username">手机:<span class="msg"></span></label>
          <input type="text" name="phone" class="am-form-field" placeholder="11纯数字手机号码" />
      </div>
     
      <div class="am-form-group form-register-sex">
      	<label class="am-radio-inline"><input type="radio" name="sex" value="1"> 男</label>
        <label class="am-radio-inline"><input type="radio" name="sex" value="0"> 女</label>
      </div>
      
      <br>
      <div class="am-cf">
        <span class="am-btn am-btn-default am-btn-primary am-btn-sm am-fl register-btn">注 册 <i class="am-icon-circle-o-notch am-icon-spin" style="display:none;"></i></span>
        <span class="am-btn am-btn-sm am-fr login-link">登 录</span>
      </div>
    </form>         
    </div>
    
  </div>
</div>
<!-- end register --> 

<div class="am-popup" id="comment-model">
  <div class="am-popup-inner">
    <div class="am-popup-hd">
      <h4 class="am-popup-title">我的评论</h4>
      <span data-am-modal-close  class="am-close">&times;</span>
    </div>
    <div class="am-popup-bd am-padding">
    
        <span id="list-mycomment-html"></span>
        <script type="text/x-handlebars-template" id="list-mycomment-template">
            <div class="am-margin-top" style="max-height:500px;">
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
                            评论于 <time datetime="">{{createtime}}</time> &nbsp; [<a href="/#blog_show-{{fid}}" target="_blank">点击查看</a>]
                            </div>
                            {{#me uid ../UID}}
                            <a href="javascript:" class="am-fr blog-comment-del-btn" data-cid="{{id}}" data-fid="{{fid}}" data-admincp=true>删除</a>
                            {{/me}}
                          </div>
                        </header>
                    
                        <div class="am-comment-bd am-text-truncate" title="{{context}}">{{context}}</div> <!-- 评论内容 -->
                      </div>
                  </li>
                  {{/each}}
                        
                    <ul class="am-pagination blog-pagination" style="text-align:center">
                    {{#if nextpage}}
                      <li class="am-pagination-next"><a href="javascript:" class=" am-text-danger" onclick="mycomment_list({{nextpage}})">下一页  »</a></li>
                    {{/if}}
                    {{#if prepage}}
                      <li class="am-pagination-next am-margin-right"><a href="javascript:" onclick="mycomment_list({{prepage}})">«  上一页</a></li>
                    {{/if}}
                    {{#compare page 1}}
                        <li class="am-pagination-next am-margin-right"><a href="javascript:" onclick="mycomment_list(1)">首页</a></li>
                    {{/compare}}
                    </ul>
                </ul>
            </div>
        </script>  
      
    </div>
  </div>
</div>
<!-- end comment-model --> 

<div class="am-popup" id="edit-model">
  <div class="am-popup-inner">
    <div class="am-popup-hd">
      <h4 class="am-popup-title">编辑信息</h4>
      <span data-am-modal-close
            class="am-close">&times;</span>
    </div>
    <div class="am-popup-bd am-padding">
        <div class="am-tabs" data-am-tabs>
          <ul class="am-tabs-nav am-nav am-nav-tabs">
            <li class="am-active"><a href="javascript: void(0)">基本信息</a></li>
            <li><a href="javascript: void(0)">修改密码</a></li>
            <li><a href="javascript: void(0)">修改邮箱</a></li>
          </ul>
          <form method="post" class="am-form" id="edit-form">
          <div class="am-tabs-bd">
            <div class="am-tab-panel am-active">
              <div class="form-register-phone am-g">
                  <div class="am-form-label am-u-sm-3">手机：</div>
                  <div class="am-u-sm-9"><input type="text" name="phone" class="am-form-field am-input-sm" placeholder="11纯数字手机号码" /></div>
              </div>
              <div class="form-register-phone am-g">
                  <div class="am-form-label am-u-sm-3">QQ：</div>
                  <div class="am-u-sm-9"><input type="text" name="qq" class="am-form-field am-input-sm" placeholder="纯数字的QQ号" /></div>
              </div>
              <div class="am-form-group am-g am-margin-top">
              	<div class="am-u-sm-3">性别：</div>
                <div class="am-u-sm-9">
                    <label class="am-radio-inline am-fl"><input type="radio" name="sex" value="1"> 男</label>
                    <label class="am-radio-inline am-fl"><input type="radio" name="sex" value="0"> 女</label>
                </div>
              </div>
              <div class="am-margin am-fr"><span class="am-btn am-btn-primary am-btn-xs edit-user-btn">提交</span></div>
            </div>
            <div class="am-tab-panel">
              <div class="form-register-phone am-g">
                  <div class="am-form-label am-u-sm-3">旧密码：</div>
                  <div class="am-u-sm-9"><input type="password" name="oldpassword" class="am-form-field am-input-sm" placeholder="输入旧密码" /></div>
              </div>
              <div class="form-register-phone am-g">
                  <div class="am-form-label am-u-sm-3">新密码：</div>
                  <div class="am-u-sm-9"><input type="password" name="newpassword" class="am-form-field am-input-sm" placeholder="6~50个字符" /></div>
              </div>
              <div class="form-register-phone am-g">
                  <div class="am-form-label am-u-sm-3">重复：</div>
                  <div class="am-u-sm-9"><input type="password" name="repassword" class="am-form-field am-input-sm" placeholder="重复新密码" /></div>
              </div>
              <div class="am-margin am-fr"><span class="am-btn am-btn-primary am-btn-xs edit-password-btn">提交</span></div>
            </div>
            <div class="am-tab-panel">
              <div class="form-register-phone am-g">
                  <div class="am-form-label am-u-sm-3">密码：</div>
                  <div class="am-u-sm-9"><input type="password" name="password" class="am-form-field am-input-sm" placeholder="输入密码" /></div>
              </div>
              <div class="form-register-phone am-g">
                  <div class="am-form-label am-u-sm-3">新邮箱：</div>
                  <div class="am-u-sm-9"><input type="text" name="email" class="am-form-field am-input-sm" placeholder="建议使用qq邮箱" /></div>
              </div>
              <div class="am-margin am-fr"><span class="am-btn am-btn-primary am-btn-xs edit-email-btn">提交</span></div>
            </div>
          </div>
        </form>
        </div>
    </div>
  </div>
</div>
<!-- end edit-model --> 

<div id="ucenter-model" class="am-offcanvas ucenter-bar"><div class="am-offcanvas-bar am-offcanvas-bar-flip" ><span id="ucenter-html"></span></div></div>
<script id="ucenter-template" type="text/x-handlebars-template">
{{#if Data.uid}}
    <div class="am-container">
        <a href="javascript:" id="avatar"><img src="{{Data.avatar}}" class="am-img-thumbnail am-circle am-center  am-margin-top" width="60" height="60" /></a>
        <div class="am-text-center am-margin-top-sm" >{{Data.username}} <img src="/static/img/sex{{Data.sex}}.png" class="sex"></div>
    </div>
         
    <div class="am-cf am-margin"></div>
    <div class="am-container am-margin-left-xs right-menu">
        <div class="am-g">
          <div class="am-u-sm-4">邮箱</div>
          <div class="am-u-sm-8 am-text-lef">{{Data.email}}</div>
        </div>
        <div class="am-g">
          <div class="am-u-sm-4">手机</div>
          <div class="am-u-sm-8 am-text-lef">{{Data.phone}}</div>
        </div>
        <div class="am-g">
          <div class="am-u-sm-4">QQ</div>
          <div class="am-u-sm-8 am-text-left">{{Data.qq}}</div>
        </div>
        <div class="am-g">
          <div class="am-u-sm-4">最后活动</div>
          <div class="am-u-sm-8 am-text-left">{{Data.updatetime}}</div>
        </div>
        <div class="am-g">
          <div class="am-u-sm-4">创建时间</div>
          <div class="am-u-sm-8 am-text-left">{{Data.createtime}}</div>
        </div>
    </div>
    <div class="ucenter-bar-line"></div>

    
    <div class="am-cf"></div>
        <ul class="am-avg-sm-2 am-text-center">
          <li class="am-padding-sm edit-link">
            <a href="javascript:"><span class="am-icon-gear am-icon-sm"></span>
            <div class="text">设置</div>
            </a>
          </li>
          <li class="am-padding-sm comment-link">
            <a href="javascript:"><span class="am-icon-edit am-icon-sm"></span>
            <div class="text">评论</div>
            </a>
          </li>
          {{#if atoken}}
          <li class="am-padding-sm">
            <a href="/admin" target="_blank"><span class="am-icon-user am-icon-sm"></span>
            <div class="text">后台管理</div>
            </a>
          </li> 
          {{/if}}         
          <li class="am-padding-sm logout-btn">
            <a href="javascript:"><span class="am-icon-sign-out am-icon-sm"></span>
            <div class="text">退出</div>
            </a>
          </li>
        </ul>           
    </div>
{{/if}}
</script>
<iframe id="iframe-form" name="iframe-form" src="" class="am-hide" ></iframe>
<form method="post" class="upload-form am-hide" enctype="multipart/form-data" target="iframe-form" action="/upload/avatar">
<input type="file" name="upload" class="upload-file hide"  />
<input type="hidden" name="uid" id="uploadUid" value="" />
</form>
<!-- end ucenter-model --> 