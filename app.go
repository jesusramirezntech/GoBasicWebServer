//
//	Author: Jesus Ramirez N
//

package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is Jesus Ramirez N!")
}

func main() {
	http.HandleFunc("/", HomeEndpoint)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}