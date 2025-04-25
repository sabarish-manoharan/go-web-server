package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported ", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello world")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/form.html")
	}

	if r.Method == "POST" {
		fmt.Fprintf(w, "POST request Successful\n")

		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprintf(w, "Name : %v\n", name)
		fmt.Fprintf(w, "Address : %v", address)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	// http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./static/form.html")
	// })
	http.HandleFunc("/hello", helloHandler)
	fmt.Print("Server is started in 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
