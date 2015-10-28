<!DOCTYPE html>

<html>
  <head>
    <title>Beego</title>
    {{template "header"}}

  </head>
<body>
	{{template "native" .}}
	<div class="container">
		<div class="page-header">
			<h1>这是我的第一篇博客</h1>
			<h6 class="text-muted">文导发表于2015年10月12日 浏览次数：200 评论次数 300</h6>
			<p>大家好。欢迎大家阅读本博客。这是一篇基础beego开发的博客,请大家多多提出意见和建议哈。如有好的建议，请发邮箱给我们</p>
		</div>
	</div>	
</body>
<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
<script src="/static/jquery.min.js"></script>

<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="/static/js/bootstrap.min.js"></script>
</html>
