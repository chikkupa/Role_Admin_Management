<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Dashboard - Bootstrap Admin Template</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
<meta name="apple-mobile-web-app-capable" content="yes">
<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
<link href="https://cdn.datatables.net/1.10.19/css/dataTables.bootstrap.min.css" rel="stylesheet">
<link href="http://fonts.googleapis.com/css?family=Open+Sans:400italic,600italic,400,600" rel="stylesheet">
<link href="/css/font-awesome.css" rel="stylesheet">
<link href="/css/style.css" rel="stylesheet">
<link href="/css/pages/dashboard.css" rel="stylesheet">
<!-- Le HTML5 shim, for IE6-8 support of HTML5 elements -->
<!--[if lt IE 9]>
      <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
</head>
<body>
<div class="navbar navbar-fixed-top">
	<div class="navbar-inner">
		<div class="container"> 
			<a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
			</a>
			<a class="brand" href="home.html">Logo <br/> <span class="wel">Welcome Admin</span></a>
			<ul class="nav pull-right">
				<li class="dropdown">
					<a href="/" class="dropdown-toggle" data-toggle="dropdown" style="text-align: right;"><i class="icon-cog"></i> Logout </a>
					<p>Last Login Time: 11:30 pm / 12-12-06-2018</p>
				</li>
			</ul>

</div></div></div>

<div class="subnavbar">
	<div class="subnavbar-inner">
		<div class="container">
			<ul class="mainnav">
				<li class="active"><a href="home.html"><span>Home</span> </a> </li>
				<li><a href="validate_user.html"><span>Validate User</span> </a> </li>
				<li><a href="secretpage.html"><span>Secret Page</span> </a></li>
			</ul>
</div></div></div>

<div class="main">
	<div class="main-inner">
		<div class="container">
			<form action="/view_user" method="post">	
				<input type="hidden" name="id" value="{{.ID}}">
			<div class="row">
				<div class="field">
					<label for="name">Name:</label>
					<label for="name">{{.Name}}</label>
				</div>	
				<div class="field">
					<label for="name">Last Name:</label>
					<label for="name">{{.LastName}}</label>
				</div>
				<div class="field">
					<label for="name">Email:</label>
					<label for="name">{{.Email}}</label>
				</div>
				<div class="field">
					<label for="name">Date of Birth:</label>
					<label for="name">{{.Dob}}</label>
				</div>	
				<div class="field">
					<label for="name">Gender:</label>
					<label for="name">{{.Gender}}</label>
				</div>
				<div class="field">
					<label for="name">Race:</label>
					<label for="name">{{.Race}}</label>
				</div>
				<div class="field">
					<label for="name">Street:</label>
					<label for="name">{{.Street}}</label>
				</div>
				<div class="field">
					<label for="name">City:</label>
					<label for="name">{{.City}}</label>
				</div>
				<div class="field">
					<label for="name">Cel:</label>
					<label for="name">{{.Phone}}</label>
				</div>
				<div class="field">
					<label for="gender">Status:</label>
					<select id="status" name="status" class="login" required>
						<option value="">Status</option>
						<option value="1" {{if eq .Status 1}}selected{{end}}>Active</option>
						<option value="2" {{if eq .Status 2}}selected{{end}}>Deactive</option>
					</select>	
				</div>
				<div style="color: red; font-weight: bold;">
					{{.Message}}
				</div>
				<div class="login-actions">
					<button class="button btn btn-primary btn-large">Register</button>
				</div>
			</div>
		</form>
</div></div></div>

<div class="footer">
	<div class="footer-inner">
		<div class="container">
			<div class="row">
				<div class="span12"> &copy; 2018 <a href="#">Admin Template</a>. </div>
</div></div></div></div>
<script src="js/jquery-1.7.2.min.js"></script> 
<script src="js/bootstrap.js"></script> 
<script src="https://cdn.datatables.net/1.10.19/js/jquery.dataTables.min.js"></script> 
<script src="https://cdn.datatables.net/1.10.19/js/dataTables.bootstrap.min.js"></script> 
<script src="js/base.js"></script> 
<script>
$(document).ready(function() {
    $('#example').DataTable();
} );
</script>
</body>
</html>
