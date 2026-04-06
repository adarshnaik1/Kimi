package http

import (
	"html/template"
	"net/http"
)

// used to panic if the err is non nil
// load all the templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter,  r * http.Request){
	err := templates.ExecuteTemplate(w,"base.html",nil)
	if err!= nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)

	}

}



