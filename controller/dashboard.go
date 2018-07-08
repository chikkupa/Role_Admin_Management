package controller

import(
    "log"
    "net/http"
    "html/template"
    "strconv"

    "../library"
    "../model"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
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