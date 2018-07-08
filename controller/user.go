package controller

import(
    "log"
    "net/http"
    "html/template"
    "strconv"
    "time"
    "math/rand"

    "../library"
    "../model"
    "../config"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:];

    if path != "" && path != "login" {
        t, _ := template.ParseFiles("views/404.html")
        t.Execute(w, nil)
    } else if r.Method == "GET" {
        error_message := ""
        t, _ := template.ParseFiles("views/login.gtpl")
        t.Execute(w, error_message)
    } else {
        r.ParseForm()
        name := r.PostFormValue("username")
        password := library.ConvertToSha1(r.PostFormValue("password"))

        user, err := model.IsValidUser(name, password)

        if err != nil {
            error_message := "Invalid username or password"
            t, _ := template.ParseFiles("views/login.gtpl")
            t.Execute(w, error_message)
            log.Print(name, ": Login Failed => Invalid credentials")
            return
        }

        library.PutCookie(w, "user_id", strconv.Itoa(user.ID))
        library.PutCookie(w, "username", user.Username)
        library.PutCookie(w, "name", user.Name)
        library.PutCookie(w, "email", user.Email)
        library.PutCookie(w, "role", strconv.Itoa(user.Role))
        library.PutCookie(w, "password", user.Password)

        // writeSuccess("Lgin Success!", w)

        log.Print(name, ": User Login Success")

        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
    }
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        user := model.User{Name: "", Username: "", Email: "", Gender: "", Message : ""}

        t, _ := template.ParseFiles("views/register.gtpl")
        t.Execute(w, user)
    } else {
        r.ParseForm()

        name := r.PostFormValue("name")
        email := r.PostFormValue("email")
        username := r.PostFormValue("username")
        password := library.ConvertToSha1(r.PostFormValue("password"))
        gender := r.PostFormValue("gender")

        user := model.User{Name: name, Username: username, Email: email, Gender: gender, Message : ""}

        errors := false

        if(!model.IsUsernameAvailable(username)){
            user.Message = "Username not available"
            errors = true
        }
        if(!errors && !model.IsEmailAvailable(email)){
            user.Message = "Email not available"
            errors = true
        }

        if(errors){
            t, _ := template.ParseFiles("views/register.gtpl")
            t.Execute(w, user)
            return
        }

        model.Register(username, name, email, password, gender)

        random_string := strconv.Itoa(rand.Intn(999)) + strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(rand.Intn(9999))
        random_string = library.ConvertToSha1(random_string)

        verify_email_link := config.BaseUrl + "/verify_email?key=" + random_string

        model.UpdateVarifyString(username, random_string)

        library.SendMail(email, "Confirm Email", "Please click the link below to verify email<br><br><a href='" + verify_email_link + "''>Verify<a>")

        data := struct{
                    Message_1 string; 
                    Message_2 string;
                }{"Your registration has been completed successfully!", "Please check your email for verification"}

        t, _ := template.ParseFiles("views/success.gtpl")
        t.Execute(w, data)
    }
}

func Verifyemail(w http.ResponseWriter, r *http.Request){
    keys, err := r.URL.Query()["key"]
    
    if !err || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    key := keys[0]

    result := model.VerifyEmail(key)

    if(result){
         data := struct{
                    Message_1 string; 
                    Message_2 template.HTML;
                }{"Your email verified successfully!", template.HTML("Please login <a href='/'>here <a>")}

        t, _ := template.ParseFiles("views/success.gtpl")
        t.Execute(w, data)
    }else{
        data := struct{
                    Message_1 string; 
                    Message_2 string;
                }{"Error verifying your email!", "Please contact the administrator for more details"}

        t, _ := template.ParseFiles("views/error.gtpl")
        t.Execute(w, data)
    }
}

func ViewUserDetails(w http.ResponseWriter, r *http.Request){
    role, _ := strconv.Atoi(library.GetCookie(r, "role"))

    if role != library.Role_admin {
        http.Redirect(w, r, "/unauthorised", http.StatusSeeOther)
    }

    if r.Method == "GET" {
        ids, err := r.URL.Query()["id"];

        if !err || len(ids[0]) < 1 {
            log.Println("Url Param 'id' is missing")
            return
        }

        id,_ := strconv.Atoi(ids[0])

        user := model.GetUserDetails(id)
        message:= library.GetCookie(r, "message")
        user.Message = message;
        library.PutCookie(w, "message", "")

        t, _ := template.ParseFiles("views/view_user.gtpl")
        t.Execute(w, user)
    }else{
        user_id, _ := strconv.Atoi(r.PostFormValue("id"))
        status, _ := strconv.Atoi(r.PostFormValue("status"))

        user := model.GetUserDetails(user_id)

        message := "";
        if user.Status != status {
            model.UpdateUserStatus(user_id, status)
            message = "User status updated successfully!"
            log.Print("Status [", user.Username, "] updated successfully")
        }else{
            message = "No change"
        }

        library.PutCookie(w, "message", message)
        http.Redirect(w, r, "/view_user?id=" + r.PostFormValue("id"), http.StatusSeeOther)
    }
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
    w.WriteHeader(status)
    w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(message))
}
