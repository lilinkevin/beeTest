<!DOCTYPE html>

<html>
  <head>
    <title>Beego</title>
    {{template "header" .}}
  </head>
<body>
	{{template "native" .}}
	<div class="container">
		<form method="POST" path="/catagory">
		  <div class="form-group">
		    <label for="title"><h4>标题</h4></label>
		    <input type="text" class="form-control" id="title" name ="title">
		  </div>
		<button type="submit" class="btn btn-default">Add New Title</button>
		</form>

	</div>	
	</br>
	<div class="container">
		<table class="table">
			<tr>
				<th>#</th>
				<th>标题</th>
				<th>创建时间</th>
				<th>访问量</th>
				<th>最后更新时间</th>
				<th>最后更新用户</th>
			</tr>
			{{range .Catagorys}}
			<tr>
				<td>#</td>
				<td>{{.Title}}</td>
				<td>{{.Ctime}}</td>
				<td>{{.Views}}</td>
				<td>{{.TopicTime}}</td>
				<td>{{.TopicUserId}}</td>
			</tr>
			{{end}}
  		</table>
	</div>
		
	
</body>
<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
<script src="/static/jquery.min.js"></script>

<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="/static/js/bootstrap.min.js"></script>
</html>
