package model

import(
    "log"
    "errors"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "../config"
)

type User struct {
    ID   int
    Name string
    LastName string
    Username string
    Gender string
    Email string
    Password string
    Role int
    Status int
}

func Register(username string, name string, email string, password string) {
	
        db, err := sql.Open(config.Mysql, config.Dbconnection)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

        stmt, err := db.Prepare("INSERT user SET username=?, name=?, email=?, password=?")
        // checkErr(err)

        res, err := stmt.Exec(username, name, email, password)

        if err != nil {
			log.Fatal(err)
		}

        id, err := res.LastInsertId()

        log.Print("New user added [Id: ", id, " Username: ", username, " Name: ", name, ", Email: ", email, "]")
}

func IsValidUser(username string, password string) (User, error){
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id, name, username, email, password, role, status from user where username=? and password=?", username, password)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    user := User{ID: 0, Name: "", Username: "", Email: "", Password: "", Role: 0, Status: 0}

    if results.Next() {
        err = results.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.Role, &user.Status)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        return user, nil;
    }

    return user, errors.New("INVALID_USERNAME_OR_PASSWORD");
}

func GetAllUsers() interface{} {
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id, name, last_name, username, gender, email, password, role, status from user")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    users := []interface{}{}

    for results.Next() {
        var user User
        // for each row, scan the result into our tag composite object
        err = results.Scan(&user.ID, &user.Name, &user.LastName, &user.Username, &user.Gender, &user.Email, &user.Password, &user.Role, &user.Status)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        users = append(users, user)
    }
    
    return users
}