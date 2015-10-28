<!DOCTYPE html>

<html>
  <head>
    <title>登陆</title>
    {{template "header"}}
  </head>
<body>
<body>
	{{template "native" .}}
	<div class="container" width="500px">

		<form method="POST" path="/login">
		  <div class="form-group">
		    <label for="Account">Account</label>
		    <input type="text" class="form-control" id="Account" name ="username" placeholder="Email">
		  </div>
		  <div class="form-group">
		    <label for="password">Password</label>
		    <input type="password" class="form-control" id="password" name ="pwd" placeholder="Password">
		  </div>
		  <div class="checkbox">
		     <label>
		       <input type="checkbox" name="autoLogin"> Auto Login
		     </label>
		   </div>
		  <button type="submit" class="btn btn-default">Submit</button>
		</form>

		
	</div>
</body>