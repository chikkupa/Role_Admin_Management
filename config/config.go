package config

var Host = "http://localhost"
var Port = "8080"

var Db_name = "go_login_app"
var Db_user = "root"
var Db_password = "123"
var Mysql = "mysql"

var Dbconnection = Db_user + ":" + Db_password + "@tcp(127.0.0.1:3306)/" + Db_name;
var Url = Host + ":" + Port