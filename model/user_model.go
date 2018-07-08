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
    Dob string
    Race string
    Email string
    Password string
    Role int
    Street string
    City string
    Phone string
    Status int
    Message string
}

func Initialize() User {
    return User{ID: 0, Name: "", LastName: "", Username: "", Gender: "", Dob: "", Race: "", Email: "", Password: "", Role: 0, Street: "", City: "", Phone: "", Status: 0,  Message: ""}
}

func Register(username string, name string, email string, password string, gender string) {
	
        db, err := sql.Open(config.Mysql, config.Dbconnection)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

        stmt, err := db.Prepare("INSERT user SET username=?, name=?, email=?, password=?, gender=?")
        // checkErr(err)

        res, err := stmt.Exec(username, name, email, password, gender)

        if err != nil {
			log.Fatal(err)
		}

        id, err := res.LastInsertId()

        log.Print("New user added [Id: ", id, " Username: ", username, " Name: ", name, ", Email: ", email, " Gender: ", gender, "]")
}

func IsValidUser(username string, password string) (User, error){
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id, name, username, email, password, role, status from user where username=? and password=? and status!=0", username, password)
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

func IsUsernameAvailable(username string) bool {
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id from user where username=? ", username)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    if results.Next() {
        return false;
    }

    return true;
}

func IsEmailAvailable(email string) bool {
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id from user where email=? ", email)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    if results.Next() {
        return false;
    }

    return true;
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

func GetUserDetails(id int) User {
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id, username, name, last_name, email, dob, gender, race, street, city, phone, role, status from user where id=?", id)

    user := Initialize()

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    if results.Next() {
        err = results.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Email, &user.Dob, &user.Gender, &user.Race, &user.Street, &user.City, &user.Phone, &user.Role, &user.Status)
   
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }

    return user;
}

func UpdateVarifyString(username string, verify_string string){
    db, err := sql.Open(config.Mysql, config.Dbconnection)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    stmt, err := db.Prepare("update user set verify_string=? where username=?")
    // checkErr(err)

    _, err = stmt.Exec(verify_string, username)

    if err != nil {
        log.Fatal(err)
    }
}

func VerifyEmail(key string) bool{
    db, err := sql.Open(config.Mysql, config.Dbconnection)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    results, err := db.Query("select id from user where verify_string=? ", key)
    if err != nil {
        return false // proper error handling instead of panic in your app
    }

    if results.Next() {
        stmt, err := db.Prepare("update user set status=2, verify_string='' where verify_string=?")
        // checkErr(err)

        _, err = stmt.Exec(key)

        if err != nil {
            log.Fatal(err)
        }
        return true
    }

    return false
}

func UpdateUserStatus(user_id int, status int){
    db, err := sql.Open(config.Mysql, config.Dbconnection)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    stmt, err := db.Prepare("update user set status=? where id=?")
    // checkErr(err)

    _, err = stmt.Exec(status, user_id)

    if err != nil {
        log.Fatal(err)
    }
}