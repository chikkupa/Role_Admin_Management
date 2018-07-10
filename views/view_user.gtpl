<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Dashboard - Bootstrap Admin Template</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
<meta name="apple-mobile-web-app-capable" content="yes">
<link href="/css/bootstrap.min.css" rel="stylesheet" type="text/css" />
<link href="http://fonts.googleapis.com/css?family=Open+Sans:400italic,600italic,400,600" rel="stylesheet">
<link href="/css/font-awesome.css" rel="stylesheet">
<link href="/css/style.css" rel="stylesheet">
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
			<div class="row">	      	
				<div class="span12">	      		
					<div class="widget">					
						<div class="widget-content">
							<div class="plan-features">
								<form action="/view_user" method="post">	
									<input type="hidden" name="id" value="{{.ID}}">
									<ul>
										<li><strong>Name: &nbsp; </strong> {{.Name}}</li>
										<li><strong>Last Name: &nbsp; </strong> {{.LastName}}</li>
										<li><strong>Email: &nbsp; </strong>  {{.Email}}</li>
										<li><strong>Date of Birth: &nbsp; </strong>  {{.Dob}}</li>
										<li><strong>Gender: &nbsp; </strong>  {{.Gender}}</li>
										<li><strong>Street: &nbsp; </strong>  {{.Street}}</li>
										<li><strong>City: &nbsp; </strong>  {{.City}}</li>
										<li><strong>Cel: &nbsp; </strong>  {{.Phone}}</li>
										<li><strong>Status: &nbsp; </strong>  
											<select id="status" name="status" class="login" required>
												<option value="">Status</option>
												<option value="1" {{if eq .Status 1}}selected{{end}}>Active</option>
												<option value="2" {{if eq .Status 2}}selected{{end}}>Deactive</option>
											</select>
										</li>
									</ul>
									<div style="color: red; font-weight: bold;">
										{{.Message}}
									</div>
									<div class="login-actions">
										<button class="button btn btn-primary btn-large">Register</button>
									</div>
								</form>	
				</div></div></div></div>
	      	
</div></div></div></div>

<div class="footer">
	<div class="footer-inner">
		<div class="container">
			<div class="row">
				<div class="span12"> &copy; 2018 <a href="#">Admin Template</a>. </div>
</div></div></div></div>
<script src="js/jquery-1.7.2.min.js"></script> 
<script src="js/bootstrap.js"></script> 
<script src="js/base.js"></script> 
</body>
</html>