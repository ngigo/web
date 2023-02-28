package main

import (
	"fmt"
	"log"
	"net/http"
)

func webserver() {
	fmt.Printf("starting Ngigo on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func login(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	fmt.Printf("login called")

	if path == "/login" {
		fmt.Printf("statement is true")
		//http.Handle("/", http.FileServer(http.Dir("./Login_v2")))
		fmt.Fprint(w, "/login is called.")
	} else {
		fmt.Fprintf(w, "Wrong path")
	}

}

func main() {
	fmt.Printf("Ngigo web starting\n")

	http.HandleFunc("/login", login)
	http.Handle("/", http.FileServer(http.Dir("./Login_v2")))
	webserver()
}
