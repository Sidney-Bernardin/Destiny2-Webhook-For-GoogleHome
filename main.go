package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"root/destinyhome"
)

var p = flag.String("p", "", "Port to listen on.")

func main() {

	// Setup router.
	r := mux.NewRouter()
	r.HandleFunc("/", destinyhome.DestinyHome).Methods("POST")

	// Listen And Serve.
	fmt.Printf("Listening on %s\n", *p)
	http.ListenAndServe(":"+*p, r)
}
