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

// this function is written to handle readiness for k8s checks.
func errors(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "you are not authorized!!", http.StatusForbidden)
}

func landingpage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, req.RemoteAddr)
}

func main() {
	fmt.Printf("Ngigo web starting\n")

	http.HandleFunc("/login", login)
	http.HandleFunc("/errors", errors)
	http.HandleFunc("/", landingpage)
	webserver()
}
