package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello %v\n", vars["id"])
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/health/check", MyHandler).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
