package main

import(
	"log"
	"net/http"

    "./config"
    "./controller"
)

func main() {
    // dbconnection = DB_USER + ":" + DB_PASSWORD + "@tcp(127.0.0.1:3306)/" + DB_NAME

    log.Print("Server startd on ", config.BaseUrl)

    http.HandleFunc("/", controller.UserLogin)
    http.HandleFunc("/login", controller.UserLogin)
    http.HandleFunc("/register", controller.UserRegister)
    http.HandleFunc("/verify_email", controller.Verifyemail)
    http.HandleFunc("/dashboard", controller.Dashboard)
    http.HandleFunc("/view_user", controller.ViewUserDetails)
    http.HandleFunc("/unauthorised", controller.CommonShowUnauthorised)
    http.HandleFunc("/show_pdf", controller.DashboarShowPdf)

    http.HandleFunc("/img/", serveResource)
    http.HandleFunc("/css/", serveResource)
    http.HandleFunc("/font/", serveResource)
    http.HandleFunc("/js/", serveResource)
    http.HandleFunc("/uploads/", serveResource)

    //added for testing
    err := http.ListenAndServe(":" + config.Port, nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func serveResource(w http.ResponseWriter, req *http.Request) {
    path := "./public" + req.URL.Path
    http.ServeFile(w, req, path)
}

