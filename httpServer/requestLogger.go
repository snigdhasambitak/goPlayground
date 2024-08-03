package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func requestLogger(w http.ResponseWriter, r *http.Request) {
	// Log the request method and URL
	fmt.Printf("Method: %s, URL: %s\n", r.Method, r.URL)

	// Log the request headers
	fmt.Println("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	// Read and log the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println("Body:")
	fmt.Println(string(body))

	// Respond to the client
	w.Write([]byte("Request received"))
}

func main() {
	http.HandleFunc("/", requestLogger)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
