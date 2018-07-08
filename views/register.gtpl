<!DOCTYPE html>
<html lang="en">
  
<head>
    <meta charset="utf-8">
    <title>Admin</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta name="apple-mobile-web-app-capable" content="yes">     
	<link href="css/bootstrap.min.css" rel="stylesheet" type="text/css" />
	<link href="css/bootstrap-responsive.min.css" rel="stylesheet" type="text/css" />

	<link href="css/font-awesome.css" rel="stylesheet">
    <link href="http://fonts.googleapis.com/css?family=Open+Sans:400italic,600italic,400,600" rel="stylesheet">
    
	<link href="css/style.css" rel="stylesheet" type="text/css">
	<link href="css/signin.css" rel="stylesheet" type="text/css">

</head>

<body>
	
<div class="account-container register">	
	<div class="content clearfix">		
		<form action="/register" method="post">		
			<h1>Signup for Free Account</h1>
			<div class="login-fields">				
				<p>Create your account:</p>
				
				<div class="field">
					<label for="name">Name:</label>
					<input type="text" id="name" name="name" value="{{.Name}}" placeholder="Name" class="login" required/>
				</div>	
				<div class="field">
					<label for="email">Email Address:</label>
					<input type="email" id="email" name="email" value="{{.Email}}" placeholder="Email" class="login" required/>
				</div>
				<div class="field">
					<label for="name">Username:</label>
					<input type="text" id="username" name="username" value="{{.Username}}" placeholder="Userame" class="login" required/>
				</div>
				<div class="field">
					<label for="password">Password:</label>
					<input type="password" id="password" name="password" value="" placeholder="Password" class="login" required/>
				</div>
				<div class="field">
					<label for="confirm_password">Confirm Password:</label>
					<input type="password" id="confirm_password" name="confirm_password" value="" placeholder="Confirm Password" class="login" required/>
					<span id="message"></span>
				</div>	
				<div class="field">
					<label for="gender">Gender:</label>
					<select id="gender" name="gender" class="login" required>
						<option value="">Gender</option>
						<option value="Male" {{if eq .Gender "Male"}}selected{{end}}>Male</option>
						<option value="Female" {{if eq .Gender "Female"}}selected{{end}}>Female</option>
					</select>	
				</div>	 
			</div>
			<div style="color: red; font-weight: bold;">
				{{.Message}}
			</div>
			<div class="login-actions">
				<button class="button btn btn-primary btn-large">Register</button>
			</div>
		</form>
</div></div>

<script src="js/jquery-1.7.2.min.js"></script>
<script src="js/bootstrap.js"></script>
<script>
	$('#confirm_password').on('keyup', function () {
		if ($('#password').val() == $('#confirm_password').val()) {
			$('#message').html('Password Matching').css('color', 'green');
		}	 else 
			$('#message').html('Password Not Matching').css('color', 'red');
	});
</script>
</body>
</html>