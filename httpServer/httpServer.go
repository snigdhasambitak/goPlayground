package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "GO app running in local")
}

func healthHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "OK\n")
}

func main(){
	fmt.Println("simple web server is starting in port : 8080")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}
