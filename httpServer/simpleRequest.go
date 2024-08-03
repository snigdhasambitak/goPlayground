package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func startServer(wg *sync.WaitGroup) {
	defer wg.Done()
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func startClient(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second) // Sleep for a second to ensure the server is up

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(body))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go startServer(&wg) // Start server in a goroutine
	go startClient(&wg) // Start client in another goroutine

	wg.Wait()
}
