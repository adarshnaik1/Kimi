package http

import (
	"net/http"
	
)

func Newrouter() *http.ServeMux{
	mux := http.NewServeMux()

	//this piece of code has been put to handle the css from the browser requests
	// browser searches for static/something.css and the file is situated at ./static/something.css
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/",HomeHandler)
	mux.HandleFunc("/post",PostHandler)
	// mux.HandleFunc("/posts",GetPostsHandler)
	
	return mux

}