package http

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct{
	Title string
	Content string
}

// used to panic if the err is non nil
// load all the templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

// Initializing a Database of Temporary Posts

var allposts = []Post{
	{Title:"The Wonderful Kimi",
	Content:"We present to you a blong aggregator made by me right here on kimi"},
}

//Handler to the HomePage
func HomeHandler(w http.ResponseWriter,  r * http.Request){
	err := templates.ExecuteTemplate(w,"base.html",nil)
	if err!= nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)

	}

}

//Handler to POST a post
func PostHandler( w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w,"Invalid Method- Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil{
		http.Error(w,"Bad Request", http.StatusBadRequest)

	}
	//Title and the Content have been figured out
	title := r.FormValue("title")
	content := r.FormValue("content")
	post := Post{Title:title, Content:content}

	allposts = append(allposts,post )
	fmt.Println("The data was appended successfully")


}


//Handler to GET posts
func GetPostsHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
		return
	}
	// Simulating the db unless we get one 
	posts:= []Post{
		{Title: "The Wonderful Kimi", Content:"We present to you a blong aggregator made by me right here on kimi"},
		{Title: "India - Israel Relations", Content: "The India Israel relations have for long be a widely debated topic among the policymakers and the security advisors"},
	}
	err := templates.ExecuteTemplate(w,"posts",posts)
	 if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}



