package main

import(
	"log"
	"net/http"
	"html/template"
	"crypto/sha1"
    "encoding/hex"
    "strconv"

    "./library"
    "./config"
    "./model"
)

func main() {
    // dbconnection = DB_USER + ":" + DB_PASSWORD + "@tcp(127.0.0.1:3306)/" + DB_NAME

    log.Print("Server startd on ", config.Url)

    http.HandleFunc("/", login)
    http.HandleFunc("/login", login)
    http.HandleFunc("/register", register)
    http.HandleFunc("/dashboard", dashboard)

    http.HandleFunc("/img/", serveResource)
    http.HandleFunc("/css/", serveResource)
    http.HandleFunc("/font/", serveResource)
    http.HandleFunc("/js/", serveResource)
    
    err := http.ListenAndServe(":" + config.Port, nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func serveResource(w http.ResponseWriter, req *http.Request) {
    path := "./public" + req.URL.Path
    http.ServeFile(w, req, path)
}

func login(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:];

    if path != "" {
        t, _ := template.ParseFiles("views/404.html")
        t.Execute(w, nil)
    } else if r.Method == "GET" {
        t, _ := template.ParseFiles("views/login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        name := r.PostFormValue("username")
        password := convertToSha1(r.PostFormValue("password"))

        user, err := model.IsValidUser(name, password)

        if err != nil {
            writeError(http.StatusBadRequest, "INVALID USERNAME OR PASSWORD", w, err)
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

func register(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
      t, _ := template.ParseFiles("views/register.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        model.Register(r.PostFormValue("username"), r.PostFormValue("name"), r.PostFormValue("email"), convertToSha1(r.PostFormValue("password")))
    }
}

func dashboard(w http.ResponseWriter, r *http.Request) {
    // name := library.GetCookie(r, "name")
    role, _ := strconv.Atoi(library.GetCookie(r, "role"))

    // writeSuccess("Hi " + name, w)
    if role == library.Role_anonymous {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }

    if role == library.Role_admin {
        users := model.GetAllUsers();
        roles := []string {"Admin", "User Type 1"}

        data := struct{
                    Roles []string; 
                    List interface{}
                }{roles, users}

        t, _ := template.ParseFiles("views/admin_dashboard.gtpl")
        t.Execute(w, data)
    }
    if role == library.Role_user_type_1 || role == library.Role_user_type_2 {
        t, _ := template.ParseFiles("views/other_dashboard.gtpl")
        t.Execute(w, nil)
    }

    log.Print("Role: ", role)
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
    w.WriteHeader(status)
    w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(message))
}

func convertToSha1(value string) string{
	h := sha1.New()
    h.Write([]byte(value))
    sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}
