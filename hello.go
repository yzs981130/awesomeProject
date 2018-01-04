package main

import (
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "hello " + message
	w.Write([]byte(message))
}

func main(){
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8080", nil)
}
