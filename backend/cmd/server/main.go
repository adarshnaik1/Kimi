package main

import (
	"fmt"
	
	"net/http"
	myrouter "github.com/adarshnaik1/go-blog/internal/http"
)

func main() {
	fmt.Printf("Application is running...")
	router := myrouter.Newrouter()
	http.ListenAndServe(":8080",router)
}