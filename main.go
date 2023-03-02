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
		fmt.Fprint(w, "you want to login? sorry but we dont know you!.")
	} else {
		fmt.Fprintf(w, "Wrong path")
	}

}

func landingpage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, req.RemoteAddr)
}

func main() {
	fmt.Printf("Ngigo web starting\n")

	http.HandleFunc("/login", login)
	http.HandleFunc("/", landingpage)
	webserver()
}
