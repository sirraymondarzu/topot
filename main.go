package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there! We are tapot."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//create a file server
	fileServer := http.FileServer(http.Dir("./static/"))

	//create url mapping to the static directory
	
	mux.Handle("/resource/", http.StripPrefix("/resource/", fileServer))
	log.Println("Starting on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
