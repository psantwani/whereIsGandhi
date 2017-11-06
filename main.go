package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	mgo "gopkg.in/mgo.v2"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	mux.HandleFunc(pat.Post("/checkin"), hello)

	http.ListenAndServe("localhost:8000", mux)
}
