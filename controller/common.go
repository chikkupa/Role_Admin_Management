package controller

import(
	"net/http"
	"html/template"
)

func CommonShowUnauthorised(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("views/unauthorised.gtpl")
    t.Execute(w, nil)
}