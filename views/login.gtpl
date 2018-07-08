<!DOCTYPE html>
<html lang="en">
  
<head>
    <meta charset="utf-8">
    <title>Admin</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta name="apple-mobile-web-app-capable" content="yes">     
	<link href="/css/bootstrap.min.css" rel="stylesheet" type="text/css" />
	<link href="/css/bootstrap-responsive.min.css" rel="stylesheet" type="text/css" />

	<link href="/css/font-awesome.css" rel="stylesheet">
    <link href="http://fonts.googleapis.com/css?family=Open+Sans:400italic,600italic,400,600" rel="stylesheet">
    
	<link href="/css/style.css" rel="stylesheet" type="text/css">
	<link href="/css/signin.css" rel="stylesheet" type="text/css">

</head>

<body>
	
<div class="account-container">	
	<img src="img/logo.png">
	<div class="content clearfix">		
		<form action="/" method="post">		
			<h1>Member Login</h1>					
			<div class="login-fields">				
				<p>Please provide your details</p>				
				<div class="field">
					<label for="username">Username</label>
					<input type="text" id="username" name="username" value="" placeholder="Username" class="login username-field" />
				</div>
				
				<div class="field">
					<label for="password">Password:</label>
					<input type="password" id="password" name="password" value="" placeholder="Password" class="login password-field"/>
				</div>				
			</div>
			<div style="color: red; font-weight: bold;">
				{{.}}
			</div>
			<div class="login-actions">		
				<a href="/register">Register</a>						
				<button class="button btn btn-success btn-large">Sign In</button>				
			</div>
		</form>
</div></div>

<script src="js/jquery-1.7.2.min.js"></script>
<script src="js/bootstrap.js"></script>
<script src="js/signin.js"></script>

</body>
</html>