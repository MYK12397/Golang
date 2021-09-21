package main

import (
	
	"io"
	"net/http"
)


type dog int

func (d dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	io.WriteString(w,"this is dog page")
}

type cat int
func (c cat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	io.WriteString(w,"this is cat page")
}


func main() {
	
	var d dog
	var c cat

	mux := http.NewServeMux()

	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)

}