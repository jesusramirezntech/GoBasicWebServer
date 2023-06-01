//
//	Author: Jesus Ramirez N
//

package main

import (
	"fmt"
	"log"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Hello, this is Jesus Ramirez N!")
}

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Jesus Ramirez N from HomeEndpoint!")
}

func main() {

	//Serving a static web page
	fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)

	fmt.Println("Starting Web Server at port 3000...")

	http.HandleFunc("/hello", hiHandler)

	http.HandleFunc("/otherport", HomeEndpoint)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}