{{define "native"}}
<nav class="navbar navbar-default">
  		<div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
   			 <div class="navbar-header">
   			 	 <a class="navbar-brand" href="#">我的博客</a>
    		</div>
			<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
     			<ul class="nav navbar-nav">
        			<li><a href="/">首页 <span class="sr-only">(current)</span></a></li>
        			<li><a href="/catagory">分类 <span class="sr-only">(current)</span></a></li>
        			<li><a href="/topic">文章<span class="sr-only">(current)</span></a></li>
  				</ul>
          <ul class="nav navbar-nav navbar-right">
            {{ if .HasLogin}}
              <li><a href="javascirpt:">欢迎您:{{.UserName}}</a></li>
              <li><a href="/logout">Logout</a></li>
            {{ else }}
              <li><a href="/login">Login</a></li>
            {{end}}

            
          </ul>
			</div>
		</div>	
	</nav>
{{end}}