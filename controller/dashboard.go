package controller

import(
    "fmt"
    "io/ioutil"
    "mime/multipart"
    "errors"
    "log"
    "net/http"
    "html/template"
    "strconv"

    "../library"
    "../model"
    "../config"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
    // name := library.GetCookie(r, "name")
    role, _ := strconv.Atoi(library.GetCookie(r, "role"))
    status, _ := strconv.Atoi(library.GetCookie(r, "status"))

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
        userDashboard(w, r)
    }

    log.Print("Role: ", role, " Status: ", status)
}

func userDashboard(w http.ResponseWriter, r *http.Request){
    user_id, _ := strconv.Atoi(library.GetCookie(r, "user_id"))
    status, _ := strconv.Atoi(library.GetCookie(r, "status"))

    if status == 2 || status == 3 {
        uploadFiles(w, r)
    }

    if status == 4 {
        files := model.GetUserFiles(user_id)
        log.Print(files)
        t, _ := template.ParseFiles("views/user_files_list.gtpl")
        t.Execute(w, files)
    }
}

// UploadFile uploads a file to the server
func uploadFiles(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        t, err := template.ParseFiles("views/upload.html")
        if err != nil {
            fmt.Println(err)
            return
        }
        t.Execute(w, nil)
        return
    }

    sUser_id := library.GetCookie(r, "user_id")
    user_id, _ := strconv.Atoi(sUser_id)

    for i := 1; i <= 8; i++ {
        index := strconv.Itoa(i)
        file_category := "field" + index
        file, handle, err := r.FormFile(file_category)

        if err != nil {
            log.Print(err.Error())
            return
        }

        filename := sUser_id + "-" + file_category + "-" +  handle.Filename
        destination := "./public/uploads/" + filename
        err = uploadPdfFile(file, handle, destination)

        if(err != nil){
            log.Print(err.Error())

            t, err := template.ParseFiles("views/upload.html")
            if err != nil {
                fmt.Println(err)
                return
            }

            t.Execute(w, err.Error())
            return
        }

        model.UserFileAdd(user_id, file_category, filename)
    }
    
    model.UpdateUserStatus(user_id, 4)
    library.PutCookie(w, "status", "4")
    log.Print("File uploaded successfully")
    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
 
func uploadImageFile(file multipart.File, handle *multipart.FileHeader, destination string) error{
    mimeType := handle.Header.Get("Content-Type")
    switch mimeType {
    case "image/jpeg", "image/jpg", "image/png":
        data, err := ioutil.ReadAll(file)
        if err != nil {
            return err
        }
        err = ioutil.WriteFile(destination, data, 0666)

        if err != nil {
            return err
        }
        return nil

    default:
        return errors.New("File type not supported")
    }
}

func uploadPdfFile(file multipart.File, handle *multipart.FileHeader, destination string) error{
    mimeType := handle.Header.Get("Content-Type")
    switch mimeType {
    case "application/pdf":
        data, err := ioutil.ReadAll(file)
        if err != nil {
            return err
        }
        err = ioutil.WriteFile(destination, data, 0666)

        if err != nil {
            return err
        }
        return nil

    default:
        return errors.New("File type not supported")
    }
}

func DashboarShowPdf(w http.ResponseWriter, r *http.Request){
    ids, err := r.URL.Query()["id"]
    
    if !err || len(ids[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

    sId := ids[0]
    id, _ := strconv.Atoi(sId)

    file := model.GetUserFileDetails(id)

    filename := config.Upload_location + file.Filename

    t, _ := template.ParseFiles("views/pdf_view.gtpl")
    t.Execute(w, filename)
}
